/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer } from 'ethers'
import { Provider, TransactionRequest } from '@ethersproject/providers'
import { Contract, ContractFactory, Overrides } from '@ethersproject/contracts'

import type { L2GatewayRouter } from '../L2GatewayRouter'

export class L2GatewayRouter__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(overrides?: Overrides): Promise<L2GatewayRouter> {
    return super.deploy(overrides || {}) as Promise<L2GatewayRouter>
  }
  getDeployTransaction(overrides?: Overrides): TransactionRequest {
    return super.getDeployTransaction(overrides || {})
  }
  attach(address: string): L2GatewayRouter {
    return super.attach(address) as L2GatewayRouter
  }
  connect(signer: Signer): L2GatewayRouter__factory {
    return super.connect(signer) as L2GatewayRouter__factory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): L2GatewayRouter {
    return new Contract(address, _abi, signerOrProvider) as L2GatewayRouter
  }
}

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'l1Token',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'gateway',
        type: 'address',
      },
    ],
    name: 'GatewaySet',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_from',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'uint256',
        name: '_transferId',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: '_amount',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'InboundTransferFinalized',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_from',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'uint256',
        name: '_transferId',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: '_amount',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'OutboundTransferInitiated',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'bool',
        name: 'success',
        type: 'bool',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_from',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: '_amount',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'bytes',
        name: 'callHookData',
        type: 'bytes',
      },
    ],
    name: 'TransferAndCallTriggered',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_userFrom',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_userTo',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'address',
        name: 'gateway',
        type: 'address',
      },
    ],
    name: 'TransferRouted',
    type: 'event',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
    ],
    name: 'calculateL2TokenAddress',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'counterpartGateway',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'defaultGateway',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_token',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_from',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_amount',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'finalizeInboundTransfer',
    outputs: [
      {
        internalType: 'bytes',
        name: '',
        type: 'bytes',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_token',
        type: 'address',
      },
    ],
    name: 'getGateway',
    outputs: [
      {
        internalType: 'address',
        name: 'gateway',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_token',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_from',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_amount',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'getOutboundCalldata',
    outputs: [
      {
        internalType: 'bytes',
        name: '',
        type: 'bytes',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_counterpartGateway',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_defaultGateway',
        type: 'address',
      },
    ],
    name: 'initialize',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    name: 'l1TokenToGateway',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_l1Token',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_amount',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'outboundTransfer',
    outputs: [
      {
        internalType: 'bytes',
        name: '',
        type: 'bytes',
      },
    ],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_token',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_amount',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: '_maxGas',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: '_gasPriceBid',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'outboundTransfer',
    outputs: [
      {
        internalType: 'bytes',
        name: '',
        type: 'bytes',
      },
    ],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'router',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'newDefaultGateway',
        type: 'address',
      },
    ],
    name: 'setDefaultGateway',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address[]',
        name: '_l1Token',
        type: 'address[]',
      },
      {
        internalType: 'address[]',
        name: '_gateway',
        type: 'address[]',
      },
    ],
    name: 'setGateway',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
]

const _bytecode =
  '0x608060405234801561001057600080fd5b50610ee0806100206000396000f3fe6080604052600436106100b25760003560e01c8063a0c76a961161006f578063a0c76a9614610410578063a7e28d48146104e9578063bda009fe1461051c578063d2ce7d651461054f578063ed08fdc6146105e9578063f7c9362f1461061c578063f887ea401461064f576100b2565b806303295802146100b75780632db09c1c146100e85780632e567b36146100fd5780634201f98514610215578063485cc955146103475780637b3a3c8b14610382575b600080fd5b3480156100c357600080fd5b506100cc610664565b604080516001600160a01b039092168252519081900360200190f35b3480156100f457600080fd5b506100cc610673565b34801561010957600080fd5b506101a0600480360360a081101561012057600080fd5b6001600160a01b03823581169260208101358216926040820135909216916060820135919081019060a081016080820135600160201b81111561016257600080fd5b82018360208201111561017457600080fd5b803590602001918460018302840111600160201b8311171561019557600080fd5b509092509050610682565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101da5781810151838201526020016101c2565b50505050905090810190601f1680156102075780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561022157600080fd5b506103456004803603604081101561023857600080fd5b810190602081018135600160201b81111561025257600080fd5b82018360208201111561026457600080fd5b803590602001918460208302840111600160201b8311171561028557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102d457600080fd5b8201836020820111156102e657600080fd5b803590602001918460208302840111600160201b8311171561030757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506106c8945050505050565b005b34801561035357600080fd5b506103456004803603604081101561036a57600080fd5b506001600160a01b0381358116916020013516610816565b6101a06004803603608081101561039857600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b8111156103d257600080fd5b8201836020820111156103e457600080fd5b803590602001918460018302840111600160201b8311171561040557600080fd5b509092509050610824565b34801561041c57600080fd5b506101a0600480360360a081101561043357600080fd5b6001600160a01b03823581169260208101358216926040820135909216916060820135919081019060a081016080820135600160201b81111561047557600080fd5b82018360208201111561048757600080fd5b803590602001918460018302840111600160201b831117156104a857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610840945050505050565b3480156104f557600080fd5b506100cc6004803603602081101561050c57600080fd5b50356001600160a01b03166108e8565b34801561052857600080fd5b506100cc6004803603602081101561053f57600080fd5b50356001600160a01b03166108f9565b6101a0600480360360c081101561056557600080fd5b6001600160a01b0382358116926020810135909116916040820135916060810135916080820135919081019060c0810160a0820135600160201b8111156105ab57600080fd5b8201836020820111156105bd57600080fd5b803590602001918460018302840111600160201b831117156105de57600080fd5b5090925090506109d0565b3480156105f557600080fd5b506100cc6004803603602081101561060c57600080fd5b50356001600160a01b0316610c65565b34801561062857600080fd5b506103456004803603602081101561063f57600080fd5b50356001600160a01b0316610c80565b34801561065b57600080fd5b506100cc610cf6565b6003546001600160a01b031681565b6000546001600160a01b031681565b6040805162461bcd60e51b815260206004820152601460248201527327a7262cafa7aaaa2127aaa7222fa927aaaa22a960611b6044820152905160609181900360640190fd5b6106d0610d05565b61071c576040805162461bcd60e51b81526020600482015260186024820152774f4e4c595f434f554e544552504152545f4741544557415960401b604482015290519081900360640190fd5b805182511461072757fe5b60005b82518110156108115781818151811061073f57fe5b60200260200101516002600085848151811061075757fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b031602179055508181815181106107af57fe5b60200260200101516001600160a01b03168382815181106107cc57fe5b60200260200101516001600160a01b03167f812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba35460405160405180910390a360010161072a565b505050565b6108208282610d16565b5050565b606061083686868660008088886109d0565b9695505050505050565b6060848260405160200180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561089c578181015183820152602001610884565b50505050905090810190601f1680156108c95780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190529a9950505050505050505050565b60006108f382610d44565b92915050565b6001600160a01b03808216600090815260026020526040902054166001811415610956576040805162461bcd60e51b815260206004820152600960248201526810931050d2d31254d560ba1b604482015290519081900360640190fd5b6001600160a01b03811661097257506003546001600160a01b03165b610984816001600160a01b0316610dd6565b6109cb576040805162461bcd60e51b81526020600482015260136024820152721393d7d1d0551155d05657d111541313d65151606a1b604482015290519081900360640190fd5b919050565b60606109da610ddc565b60006109e5896108f9565b90506060610a2c8a338b8b89898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061084092505050565b604080516001600160a01b0385811682529151929350818c169233928e16917f85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5919081900360200190a4816001600160a01b031663d2ce7d65348c8c8c8c8c886040518863ffffffff1660e01b815260040180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b0316815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610b1a578181015183820152602001610b02565b50505050905090810190601f168015610b475780820380516001836020036101000a031916815260200191505b509750505050505050506000604051808303818588803b158015610b6a57600080fd5b505af1158015610b7e573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f191682016040526020811015610ba857600080fd5b8101908080516040519392919084600160201b821115610bc757600080fd5b908301906020820185811115610bdc57600080fd5b8251600160201b811182820188101715610bf557600080fd5b82525081516020918201929091019080838360005b83811015610c22578181015183820152602001610c0a565b50505050905090810190601f168015610c4f5780820380516001836020036101000a031916815260200191505b5060405250505092505050979650505050505050565b6002602052600090815260409020546001600160a01b031681565b610c88610d05565b610cd4576040805162461bcd60e51b81526020600482015260186024820152774f4e4c595f434f554e544552504152545f4741544557415960401b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b6000546001600160a01b0316331490565b610d21826000610dde565b600380546001600160a01b0319166001600160a01b039290921691909117905550565b6000610d4f826108f9565b6001600160a01b031663a7e28d48836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610da457600080fd5b505afa158015610db8573d6000803e3d6000fd5b505050506040513d6020811015610dce57600080fd5b505192915050565b3b151590565b565b6001600160a01b038216610e2f576040805162461bcd60e51b81526020600482015260136024820152721253959053125117d0d3d55395115494105495606a1b604482015290519081900360640190fd5b6000546001600160a01b031615610e7c576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600080546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905556fea26469706673582212206e98a18d5551a485a4e1ba53c046c88088c0f94f7d0a1e4cdf7f40cc706c2c7164736f6c634300060b0033'
