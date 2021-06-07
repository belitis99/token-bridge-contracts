/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer } from 'ethers'
import { Provider, TransactionRequest } from '@ethersproject/providers'
import { Contract, ContractFactory, Overrides } from '@ethersproject/contracts'

import type { L2GatewayTester } from '../L2GatewayTester'

export class L2GatewayTester__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(overrides?: Overrides): Promise<L2GatewayTester> {
    return super.deploy(overrides || {}) as Promise<L2GatewayTester>
  }
  getDeployTransaction(overrides?: Overrides): TransactionRequest {
    return super.getDeployTransaction(overrides || {})
  }
  attach(address: string): L2GatewayTester {
    return super.attach(address) as L2GatewayTester
  }
  connect(signer: Signer): L2GatewayTester__factory {
    return super.connect(signer) as L2GatewayTester__factory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): L2GatewayTester {
    return new Contract(address, _abi, signerOrProvider) as L2GatewayTester
  }
}

const _abi = [
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
    inputs: [],
    name: 'beacon',
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
    name: 'cloneableProxyHash',
    outputs: [
      {
        internalType: 'bytes32',
        name: '',
        type: 'bytes32',
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
    name: 'exitNum',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
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
        name: 'outboundCalldata',
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
        name: '_l1Counterpart',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_router',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_beacon',
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
        internalType: 'contract IArbToken',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        internalType: 'address',
        name: 'sender',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'dest',
        type: 'address',
      },
      {
        internalType: 'bytes',
        name: 'data',
        type: 'bytes',
      },
    ],
    name: 'mintAndCall',
    outputs: [],
    stateMutability: 'nonpayable',
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
]

const _bytecode =
  '0x608060405234801561001057600080fd5b506121ff806100206000396000f3fe608060405260043610620000b65760003560e01c8063a0c76a96116200006d578063a0c76a9614620002ff578063a7e28d4814620003df578063c0c53b8b1462000416578063d2ce7d651462000461578063db70cf6e1462000501578063f887ea4014620005df57620000b6565b8063015234ab14620000bb5780632db09c1c14620000e55780632e567b36146200011957806359659e90146200023b5780637b3a3c8b146200025357806397881f8d14620002e7575b600080fd5b348015620000c857600080fd5b50620000d3620005f7565b60408051918252519081900360200190f35b348015620000f257600080fd5b50620000fd620005fd565b604080516001600160a01b039092168252519081900360200190f35b3480156200012657600080fd5b50620001c3600480360360a08110156200013f57600080fd5b6001600160a01b03823581169260208101358216926040820135909216916060820135919081019060a081016080820135600160201b8111156200018257600080fd5b8201836020820111156200019557600080fd5b803590602001918460018302840111600160201b83111715620001b757600080fd5b5090925090506200060c565b6040805160208082528351818301528351919283929083019185019080838360005b83811015620001ff578181015183820152602001620001e5565b50505050905090810190601f1680156200022d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156200024857600080fd5b50620000fd62000b9e565b620001c3600480360360808110156200026b57600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b811115620002a657600080fd5b820183602082011115620002b957600080fd5b803590602001918460018302840111600160201b83111715620002db57600080fd5b50909250905062000bad565b348015620002f457600080fd5b50620000d362000bc1565b3480156200030c57600080fd5b50620001c3600480360360a08110156200032557600080fd5b6001600160a01b03823581169260208101358216926040820135909216916060820135919081019060a081016080820135600160201b8111156200036857600080fd5b8201836020820111156200037b57600080fd5b803590602001918460018302840111600160201b831117156200039d57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955062000bf0945050505050565b348015620003ec57600080fd5b50620000fd600480360360208110156200040557600080fd5b50356001600160a01b031662000d6d565b3480156200042357600080fd5b506200045f600480360360608110156200043c57600080fd5b506001600160a01b03813581169160208101358216916040909101351662000dca565b005b620001c3600480360360c08110156200047957600080fd5b6001600160a01b0382358116926020810135909116916040820135916060810135916080820135919081019060c0810160a0820135600160201b811115620004c057600080fd5b820183602082011115620004d357600080fd5b803590602001918460018302840111600160201b83111715620004f557600080fd5b50909250905062000e95565b3480156200050e57600080fd5b506200045f600480360360a08110156200052757600080fd5b6001600160a01b0382358116926020810135926040820135831692606083013516919081019060a081016080820135600160201b8111156200056857600080fd5b8201836020820111156200057b57600080fd5b803590602001918460018302840111600160201b831117156200059d57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955062001134945050505050565b348015620005ec57600080fd5b50620000fd620013aa565b60025481565b6000546001600160a01b031681565b606062000618620013b9565b62000665576040805162461bcd60e51b81526020600482015260186024820152774f4e4c595f434f554e544552504152545f4741544557415960401b604482015290519081900360640190fd5b606080848460408110156200067957600080fd5b810190602081018135600160201b8111156200069457600080fd5b820183602082011115620006a757600080fd5b803590602001918460018302840111600160201b83111715620006c957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156200071c57600080fd5b8201836020820111156200072f57600080fd5b803590602001918460018302840111600160201b831117156200075157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509799509297506200079d95508f9450620013be9350505050565b9050620007b3816001600160a01b03166200140e565b620007ee576000620007c78b838662001414565b90508015620007ec576040518060200160405280600081525094505050505062000b94565b505b815181901562000a62576000306001600160a01b031663db70cf6e838b8e8e896040518663ffffffff1660e01b815260040180866001600160a01b03166001600160a01b03168152602001858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015620008aa57818101518382015260200162000890565b50505050905090810190601f168015620008d85780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b158015620008fc57600080fd5b505af19250505080156200090e575060015b6200099357816001600160a01b0316638c2a993e8c8b6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156200097457600080fd5b505af115801562000989573d6000803e3d6000fd5b5050505062000997565b5060015b896001600160a01b03168b6001600160a01b03167f11ff8525c5d96036231ee652c108808dee4c40728a6117830a75029298bb7de6838c88604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101562000a1e57818101518382015260200162000a04565b50505050905090810190601f16801562000a4c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35062000add565b806001600160a01b0316638c2a993e8a8a6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801562000ac357600080fd5b505af115801562000ad8573d6000803e3d6000fd5b505050505b816001600160a01b0316896001600160a01b03168b6001600160a01b03167f179a84706122b1b806f7d61c28c5caef276b7ff474ae596df3cad4bbaf0eb97d8e8c8c8c60405180856001600160a01b03166001600160a01b03168152602001848152602001806020018281038252848482818152602001925080828437600083820152604051601f909101601f191690920182900397509095505050505050a4604051806020016040528060008152509450505050505b9695505050505050565b6003546001600160a01b031681565b606062000b94868686600080888862000e95565b60405162000bd2602082016200192f565b6020820181038252601f19601f820116604052508051906020012081565b6060632e567b3660e01b86868686600254876040516020018083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101562000c4a57818101518382015260200162000c30565b50505050905090810190601f16801562000c785780820380516001836020036101000a031916815260200191505b5060408051601f19818403018152908290526001600160a01b03808b16602484019081528a8216604485015290891660648401526084830188905260a060a48401908152825160c48501528251929750909550935060e49091019150602085019080838360005b8381101562000cf957818101518382015260200162000cdf565b50505050905090810190601f16801562000d275780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909a16999099179098525095965050505050505095945050505050565b600062000d79620015ab565b62000db9576040805162461bcd60e51b815260206004820152600b60248201526a27a7262cafa927aaaa22a960a91b604482015290519081900360640190fd5b62000dc482620013be565b92915050565b62000dd68383620015bc565b6001600160a01b03811662000e23576040805162461bcd60e51b815260206004820152600e60248201526d24a72b20a624a22fa122a0a1a7a760911b604482015290519081900360640190fd5b6003546001600160a01b03161562000e71576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b03929092169190911790555050565b6060341562000ed6576040805162461bcd60e51b81526020600482015260086024820152674e4f5f56414c554560c01b604482015290519081900360640190fd5b6000606062000f1b85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250620015cc92505050565b9150915060008062000f2d8c620013be565b905062000f43816001600160a01b03166200140e565b62000f8a576040805162461bcd60e51b81526020600482015260126024820152711513d2d15397d393d517d111541313d6515160721b604482015290519081900360640190fd5b806001600160a01b03166374f4f547858c6040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801562000feb57600080fd5b505af115801562001000573d6000803e3d6000fd5b505050506060620010158d868e8e8862000bf0565b90506200102281620016c9565b92505050600260008154809291906001019190505550808a6001600160a01b0316846001600160a01b03167f9c003a9d1163eca79021710dcd5d9f657218bf2bd67aaa13389009a6047894a88e8d8760405180846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015620010ce578181015183820152602001620010b4565b50505050905090810190601f168015620010fc5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a46040805160208082019390935281518082039093018352810190529a9950505050505050505050565b33301462001189576040805162461bcd60e51b815260206004820152601f60248201527f4d696e742063616e206f6e6c792062652063616c6c65642062792073656c6600604482015290519081900360640190fd5b6200119d826001600160a01b03166200140e565b620011ef576040805162461bcd60e51b815260206004820152601e60248201527f44657374696e6174696f6e206d757374206265206120636f6e74726163740000604482015290519081900360640190fd5b846001600160a01b0316638c2a993e83866040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156200125057600080fd5b505af115801562001265573d6000803e3d6000fd5b50505050600062001275620017d0565b5a039050805a11620012b95760405162461bcd60e51b815260040180806020018281038252602b8152602001806200219f602b913960400191505060405180910390fd5b826001600160a01b031663a4c0ed36828688866040518563ffffffff1660e01b815260040180846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015620013385781810151838201526020016200131e565b50505050905090810190601f168015620013665780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600088803b1580156200138857600080fd5b5087f11580156200139d573d6000803e3d6000fd5b5050505050505050505050565b6001546001600160a01b031681565b600190565b600080620013cc83620017d6565b9050620014078160405180602001620013e5906200192f565b6020820181038252601f19601f82011660405250805190602001203062001806565b9392505050565b3b151590565b6000806200142285620017d6565b905060008160405162001435906200192f565b8190604051809103906000f590508015801562001456573d6000803e3d6000fd5b509050806001600160a01b031663189db7d287866040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015620014d0578181015183820152602001620014b6565b50505050905090810190601f168015620014fe5780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b1580156200151f57600080fd5b505af115801562001534573d6000803e3d6000fd5b50505050846001600160a01b0316816001600160a01b031614156200155d5760009250620015a2565b6040805162461bcd60e51b81526020600482015260156024820152742ba927a723afa222a82627aca6a2a72a2fa0a2222960591b604482015290519081900360640190fd5b50509392505050565b6001546001600160a01b0316331490565b620015c8828262001861565b5050565b60006060620015da620015ab565b15620016be57828060200190516040811015620015f657600080fd5b815160208301805160405192949293830192919084600160201b8211156200161d57600080fd5b9083019060208201858111156200163357600080fd5b8251600160201b8111828201881017156200164d57600080fd5b82525081516020918201929091019080838360005b838110156200167c57818101518382015260200162001662565b50505050905090810190601f168015620016aa5780820380516001836020036101000a031916815260200191505b5060405250929450909250620016c4915050565b50339050815b915091565b60008054604051835183926060926001600160a01b0390911691869190819060208401908083835b60208310620017125780518252601f199092019160209182019101620016f1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811462001776576040519150601f19603f3d011682016040523d82523d6000602084013e6200177b565b606091505b509150915081620017c5576040805162461bcd60e51b815260206004820152600f60248201526e13d5551093d5539117d49155915495608a1b604482015290519081900360640190fd5b506105399392505050565b61c35090565b604080516001600160a01b0383166020808301919091528251808303820181529183019092528051910120919050565b604080516001600160f81b031960208083019190915260609390931b6bffffffffffffffffffffffff191660218201526035810194909452605580850193909352805180850390930183526075909301909252805191012090565b6001600160a01b038216620018b3576040805162461bcd60e51b81526020600482015260136024820152721253959053125117d0d3d55395115494105495606a1b604482015290519081900360640190fd5b6000546001600160a01b03161562001901576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b600080546001600160a01b039384166001600160a01b03199182161790915560018054929093169116179055565b610861806200193e8339019056fe608060405234801561001057600080fd5b50336001600160a01b03166359659e906040518163ffffffff1660e01b815260040160206040518083038186803b15801561004a57600080fd5b505afa15801561005e573d6000803e3d6000fd5b505050506040513d602081101561007457600080fd5b5051604080516020810182526000815281517f656970313936372e70726f78792e626561636f6e00000000000000000000000081529151918290036014019091206000805160206107a1833981519152600019909101146100d157fe5b6100e482826001600160e01b036100eb16565b505061049a565b6100fe8261024960201b6100311760201c565b6101395760405162461bcd60e51b81526004018080602001828103825260258152602001806107e26025913960400191505060405180910390fd5b6101b1826001600160a01b0316635c60da1b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561017557600080fd5b505afa158015610189573d6000803e3d6000fd5b505050506040513d602081101561019f57600080fd5b5051610249602090811b61003117901c565b6101ec5760405162461bcd60e51b815260040180806020018281038252603481526020018061082d6034913960400191505060405180910390fd5b6000805160206107a1833981519152828155815115610244576102426102196001600160e01b0361024f16565b836040518060600160405280602181526020016107c1602191396102cb60201b6100371760201c565b505b505050565b3b151590565b60006102626001600160e01b036103e316565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561029a57600080fd5b505afa1580156102ae573d6000803e3d6000fd5b505050506040513d60208110156102c457600080fd5b5051905090565b60606102df846001600160e01b0361024916565b61031a5760405162461bcd60e51b81526004018080602001828103825260268152602001806108076026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106103585780518252601f199092019160209182019101610339565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146103b8576040519150601f19603f3d011682016040523d82523d6000602084013e6103bd565b606091505b5090925090506103d78282866001600160e01b036103f616565b925050505b9392505050565b6000805160206107a18339815191525490565b606083156104055750816103dc565b8251156104155782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561045f578181015183820152602001610447565b50505050905090810190601f16801561048c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b6102f8806104a96000396000f3fe60806040523661001357610011610017565b005b6100115b61001f61002f565b61002f61002a61013c565b6101af565b565b3b151590565b606061004284610031565b61007d5760405162461bcd60e51b815260040180806020018281038252602681526020018061029d6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106100bb5780518252601f19909201916020918201910161009c565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461011b576040519150601f19603f3d011682016040523d82523d6000602084013e610120565b606091505b50915091506101308282866101d3565b925050505b9392505050565b6000610146610277565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561017e57600080fd5b505afa158015610192573d6000803e3d6000fd5b505050506040513d60208110156101a857600080fd5b5051905090565b3660008037600080366000845af43d6000803e8080156101ce573d6000f35b3d6000fd5b606083156101e2575081610135565b8251156101f25782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561023c578181015183820152602001610224565b50505050905090810190601f1680156102695780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50549056fe416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e7472616374a26469706673582212201d76e40c8b547001f2679ed8bea580bcbb5d716c3d42ba31a888159329d0181364736f6c634300060b0033a3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50426561636f6e50726f78793a2066756e6374696f6e2063616c6c206661696c6564426561636f6e50726f78793a20626561636f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e7472616374426561636f6e50726f78793a20626561636f6e20696d706c656d656e746174696f6e206973206e6f74206120636f6e74726163744d696e7420616e642063616c6c20676173206c6566742063616c63756c6174696f6e20756e6465666c6f77a2646970667358221220bba71249e8733cef5e8cd7422aedf272ac16126859313229a5af62236f5d2b1c64736f6c634300060b0033'
