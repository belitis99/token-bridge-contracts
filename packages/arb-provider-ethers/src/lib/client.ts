import * as ArbValue from './value';

import * as ethers from 'ethers';

const jaysonBrowserClient = require('jayson/lib/client/browser');
const fetch = require('node-fetch');

export enum EVMCode {
    Revert = 0,
    Invalid = 1,
    Return = 2,
    Stop = 3,
    BadSequenceCode = 4,
}

function logValToLog(val: ArbValue.Value) {
    const value = val as ArbValue.TupleValue;
    return {
        contractId: (value.get(0) as ArbValue.IntValue).bignum,
        data: ArbValue.sizedByteRangeToBytes(value.get(1) as ArbValue.TupleValue),
        topics: value.contents.slice(2).map(rawTopic => (rawTopic as ArbValue.IntValue).bignum),
    };
}

function stackValueToList(value: ArbValue.TupleValue) {
    const values = [];
    while (value.contents.length !== 0) {
        values.push(value.get(1));
        value = value.get(0) as ArbValue.TupleValue;
    }
    return values;
}

class OrigMessage {
    public data: Uint8Array;
    public calldataHash: string;
    public contractID: string;
    public sequenceNum: string;
    public timestamp: string;
    public blockHeight: string;
    public txHash: string;
    public tokenType: string;
    public value: string;
    public caller: string;

    constructor(value: ArbValue.TupleValue) {
        const wrappedData = value.get(0) as ArbValue.TupleValue;
        const calldata = wrappedData.get(0) as ArbValue.TupleValue;
        this.calldataHash = calldata.hash();
        this.data = ArbValue.sizedByteRangeToBytes(calldata.get(0) as ArbValue.TupleValue);
        this.contractID = ethers.utils.getAddress((calldata.get(1) as ArbValue.IntValue).bignum.toHexString());
        this.sequenceNum = (calldata.get(2) as ArbValue.IntValue).bignum.toHexString();
        this.timestamp = (wrappedData.get(1) as ArbValue.IntValue).bignum.toHexString();
        this.blockHeight = (wrappedData.get(2) as ArbValue.IntValue).bignum.toHexString();
        this.txHash = (wrappedData.get(3) as ArbValue.IntValue).bignum.toHexString();
        this.tokenType = (value.get(3) as ArbValue.IntValue).bignum.toHexString();
        this.value = (value.get(2) as ArbValue.IntValue).bignum.toHexString();
        this.caller = ethers.utils.getAddress((value.get(1) as ArbValue.IntValue).bignum.toHexString());
    }
}

type EVMResult = EVMReturn | EVMRevert | EVMStop | EVMBadSequenceCode | EVMInvalid;

export class EVMReturn {
    public orig: OrigMessage;
    public data: Uint8Array;
    public logs: any;
    public returnType: EVMCode.Return;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.data = ArbValue.sizedByteRangeToBytes(value.get(2) as ArbValue.TupleValue);
        this.logs = stackValueToList(value.get(1) as ArbValue.TupleValue).map(logValToLog);
        this.returnType = EVMCode.Return;
    }
}

export class EVMRevert {
    public orig: OrigMessage;
    public data: Uint8Array;
    public returnType: EVMCode.Revert;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.data = ArbValue.sizedByteRangeToBytes(value.get(2) as ArbValue.TupleValue);
        this.returnType = EVMCode.Revert;
    }
}

export class EVMStop {
    public orig: OrigMessage;
    public logs: any;
    public returnType: EVMCode.Stop;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.logs = stackValueToList(value.get(1) as ArbValue.TupleValue).map(logValToLog);
        this.returnType = EVMCode.Stop;
    }
}

export class EVMBadSequenceCode {
    public orig: OrigMessage;
    public returnType: EVMCode.BadSequenceCode;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.returnType = EVMCode.BadSequenceCode;
    }
}

export class EVMInvalid {
    public orig: OrigMessage;
    public returnType: EVMCode.Invalid;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.returnType = EVMCode.Invalid;
    }
}

function processLog(value: ArbValue.TupleValue): EVMResult {
    const returnCode = value.get(3) as ArbValue.IntValue;
    switch (returnCode.bignum.toNumber()) {
        case EVMCode.Return:
            return new EVMReturn(value);
        case EVMCode.Revert:
            return new EVMRevert(value);
        case EVMCode.Stop:
            return new EVMStop(value);
        case EVMCode.BadSequenceCode:
            return new EVMBadSequenceCode(value);
        case EVMCode.Invalid:
            return new EVMInvalid(value);
        default:
            throw Error('processLogs Invalid EVM return code');
    }
}

interface IGetVMInfoReply {
    vmID: string;
}

interface IGetAssertionCountReply {
    assertionCount: number;
}

interface IGetMessageResultReply {
    found: boolean;
    rawVal: string;
    logPreHash: string;
    logPostHash: string;
    logValHashes: string[];
    validatorSigs: string[];
    partialHash: string;
    onChainTxHash: string;
}

interface ISendMessageReply {
    txHash: string;
}

interface ICallMessageReply {
    rawVal: string;
}

interface ILogInfo {
    address: string;
    blockHash: string;
    blockNumber: string;
    data: string;
    logIndex: string;
    topics: string[];
    transactionIndex: string;
    transactionHash: string;
}

interface IFindLogsReply {
    logs: ILogInfo[];
}

function _arbClient(managerAddress: string) {
    const callServer = (request: any, callback: any) => {
        const options = {
            body: request, // request is a string
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST',
        };

        fetch(managerAddress, options)
            .then((res: any) => {
                return res.text();
            })
            .then((text: any) => {
                callback(null, text);
            })
            .catch((err: any) => {
                callback(err);
            });
    };

    return jaysonBrowserClient(callServer, {});
}

export class ArbClient {
    public client: any;

    constructor(managerUrl: string) {
        this.client = _arbClient(managerUrl);
    }

    public async getMessageResult(txHash: string) {
        const messageResult = await new Promise<IGetMessageResultReply>((resolve, reject) => {
            this.client.request(
                'Validator.GetMessageResult',
                [
                    {
                        txHash,
                    },
                ],
                (err: Error, error: Error, result: IGetMessageResultReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result);
                    }
                },
            );
        });
        if (messageResult.found) {
            const vmId = await this.getVmID();
            const val = ArbValue.unmarshal(messageResult.rawVal);
            const evmVal = processLog(val as ArbValue.TupleValue);
            let logValHashes = messageResult.logValHashes;
            if (!logValHashes) {
                logValHashes = [];
            }

            const data = {
                logPostHash: messageResult.logPostHash,
                logPreHash: messageResult.logPreHash,
                logValHashes,
                onChainTxHash: messageResult.onChainTxHash,
                partialHash: messageResult.partialHash,
                val,
                validatorSigs: messageResult.validatorSigs,
                vmId,
            };

            return {
                data,
                evmVal,
            };
        } else {
            return null;
        }
    }

    public sendMessage(value: ArbValue.Value, sig: string, pubkey: string): Promise<string> {
        return new Promise((resolve, reject) => {
            this.client.request(
                'Validator.SendMessage',
                [
                    {
                        data: ethers.utils.hexlify(ArbValue.marshal(value)),
                        pubkey,
                        signature: sig,
                    },
                ],
                (err: Error, error: Error, result: ISendMessageReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result.txHash);
                    }
                },
            );
        });
    }

    public call(value: ArbValue.Value, sender: string): Promise<Uint8Array> {
        return new Promise((resolve, reject) => {
            this.client.request(
                'Validator.CallMessage',
                [
                    {
                        data: ethers.utils.hexlify(ArbValue.marshal(value)),
                        sender,
                    },
                ],
                (err: Error, error: Error, result: ICallMessageReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        const val = ArbValue.unmarshal(result.rawVal);
                        const evmVal = processLog(val as ArbValue.TupleValue);
                        switch (evmVal.returnType) {
                            case EVMCode.Return:
                                resolve(evmVal.data);
                                break;
                            case EVMCode.Stop:
                                resolve(new Uint8Array());
                                break;
                            default:
                                reject(new Error('Call was reverted'));
                                break;
                        }
                    }
                },
            );
        });
    }

    public findLogs(fromBlock: number, toBlock: number, address: string, topics: string[]): Promise<ILogInfo[]> {
        return new Promise((resolve, reject) => {
            return this.client.request(
                'Validator.FindLogs',
                [
                    {
                        address,
                        fromHeight: fromBlock,
                        toHeight: toBlock,
                        topics,
                    },
                ],
                (err: Error, error: Error, result: IFindLogsReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result.logs);
                    }
                },
            );
        });
    }

    public getVmID(): Promise<string> {
        return new Promise((resolve, reject) => {
            this.client.request('Validator.GetVMInfo', [], (err: Error, error: Error, result: IGetVMInfoReply) => {
                if (err) {
                    reject(err);
                } else if (error) {
                    reject(error);
                } else {
                    resolve(result.vmID);
                }
            });
        });
    }

    public getAssertionCount(): Promise<number> {
        return new Promise((resolve, reject) => {
            this.client.request(
                'Validator.GetAssertionCount',
                [],
                (err: Error, error: Error, result: IGetAssertionCountReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result.assertionCount);
                    }
                },
            );
        });
    }

    public getVMCreatedTxHash(): Promise<string> {
        return new Promise((resolve, reject) => {
            this.client.request('Validator.GetVMCreatedTxHash', [], (err: Error, error: Error, result: any) => {
                if (err) {
                    reject(err);
                } else if (error) {
                    reject(error);
                } else {
                    resolve(result.vmCreatedTxHash);
                }
            });
        });
    }
}
