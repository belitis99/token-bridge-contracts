/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer } from 'ethers'
import { Provider, TransactionRequest } from '@ethersproject/providers'
import { Contract, ContractFactory, Overrides } from '@ethersproject/contracts'

import type { TestArbCustomToken } from '../TestArbCustomToken'

export class TestArbCustomToken__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(
    _bridge: string,
    _l1Address: string,
    overrides?: Overrides
  ): Promise<TestArbCustomToken> {
    return super.deploy(
      _bridge,
      _l1Address,
      overrides || {}
    ) as Promise<TestArbCustomToken>
  }
  getDeployTransaction(
    _bridge: string,
    _l1Address: string,
    overrides?: Overrides
  ): TransactionRequest {
    return super.getDeployTransaction(_bridge, _l1Address, overrides || {})
  }
  attach(address: string): TestArbCustomToken {
    return super.attach(address) as TestArbCustomToken
  }
  connect(signer: Signer): TestArbCustomToken__factory {
    return super.connect(signer) as TestArbCustomToken__factory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): TestArbCustomToken {
    return new Contract(address, _abi, signerOrProvider) as TestArbCustomToken
  }
}

const _abi = [
  {
    inputs: [
      {
        internalType: 'address',
        name: '_bridge',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_l1Address',
        type: 'address',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'owner',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'spender',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'value',
        type: 'uint256',
      },
    ],
    name: 'Approval',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'value',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'bytes',
        name: 'data',
        type: 'bytes',
      },
    ],
    name: 'Transfer',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'value',
        type: 'uint256',
      },
    ],
    name: 'Transfer',
    type: 'event',
  },
  {
    inputs: [],
    name: 'DOMAIN_SEPARATOR',
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
    inputs: [
      {
        internalType: 'address',
        name: 'owner',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'spender',
        type: 'address',
      },
    ],
    name: 'allowance',
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
        name: 'spender',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'approve',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'account',
        type: 'address',
      },
    ],
    name: 'balanceOf',
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
    inputs: [],
    name: 'bridge',
    outputs: [
      {
        internalType: 'contract ArbTokenBridge',
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
        name: 'account',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'bridgeBurn',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'account',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'bridgeMint',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'decimals',
    outputs: [
      {
        internalType: 'uint8',
        name: '',
        type: 'uint8',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'spender',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'subtractedValue',
        type: 'uint256',
      },
    ],
    name: 'decreaseAllowance',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'spender',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'addedValue',
        type: 'uint256',
      },
    ],
    name: 'increaseAllowance',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'string',
        name: 'name',
        type: 'string',
      },
      {
        internalType: 'string',
        name: 'symbol',
        type: 'string',
      },
      {
        internalType: 'uint8',
        name: 'decimals',
        type: 'uint8',
      },
    ],
    name: 'initialize',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'l1Address',
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
    name: 'name',
    outputs: [
      {
        internalType: 'string',
        name: '',
        type: 'string',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'owner',
        type: 'address',
      },
    ],
    name: 'nonces',
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
        name: 'owner',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'spender',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'value',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: 'deadline',
        type: 'uint256',
      },
      {
        internalType: 'uint8',
        name: 'v',
        type: 'uint8',
      },
      {
        internalType: 'bytes32',
        name: 'r',
        type: 'bytes32',
      },
      {
        internalType: 'bytes32',
        name: 's',
        type: 'bytes32',
      },
    ],
    name: 'permit',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'someWackyCustomStuff',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'symbol',
    outputs: [
      {
        internalType: 'string',
        name: '',
        type: 'string',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'totalSupply',
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
        name: 'recipient',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'transfer',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_value',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'transferAndCall',
    outputs: [
      {
        internalType: 'bool',
        name: 'success',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'sender',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'recipient',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'transferFrom',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
]

const _bytecode =
  '0x60806040523480156200001157600080fd5b506040516200264638038062002646833981810160405260408110156200003757600080fd5b50805160209182015160cc80546001600160a01b038085166001600160a01b03199283161790925560cd805492841692909116919091179055604080518082018252600f81526e2a32b9ba21bab9ba37b6aa37b5b2b760891b818601528151808301909252600482526321a0a92160e11b8286015292939192620000ca92909190601290620006ef620000d2821b17901c565b5050620007c5565b600054610100900460ff1680620000f75750620000f76001600160e01b03620001c516565b8062000106575060005460ff16155b620001435760405162461bcd60e51b815260040180806020018281038252602e81526020018062002618602e913960400191505060405180910390fd5b600054610100900460ff161580156200016f576000805460ff1961ff0019909116610100171660011790555b62000183846001600160e01b03620001e416565b6200019884846001600160e01b03620002ea16565b620001ac826001600160e01b03620003c716565b8015620001bf576000805461ff00191690555b50505050565b6000620001dd30620003dd60201b62000d441760201c565b1590505b90565b600054610100900460ff1680620002095750620002096001600160e01b03620001c516565b8062000218575060005460ff16155b620002555760405162461bcd60e51b815260040180806020018281038252602e81526020018062002618602e913960400191505060405180910390fd5b600054610100900460ff1615801562000281576000805460ff1961ff0019909116610100171660011790555b620002946001600160e01b03620003e316565b620002bf82604051806040016040528060018152602001603160f81b8152506200049660201b60201c565b620002d3826001600160e01b036200056716565b8015620002e6576000805461ff00191690555b5050565b600054610100900460ff16806200030f57506200030f6001600160e01b03620001c516565b806200031e575060005460ff16155b6200035b5760405162461bcd60e51b815260040180806020018281038252602e81526020018062002618602e913960400191505060405180910390fd5b600054610100900460ff1615801562000387576000805460ff1961ff0019909116610100171660011790555b6200039a6001600160e01b03620003e316565b620003af83836001600160e01b036200063616565b8015620003c2576000805461ff00191690555b505050565b6038805460ff191660ff92909216919091179055565b3b151590565b600054610100900460ff1680620004085750620004086001600160e01b03620001c516565b8062000417575060005460ff16155b620004545760405162461bcd60e51b815260040180806020018281038252602e81526020018062002618602e913960400191505060405180910390fd5b600054610100900460ff1615801562000480576000805460ff1961ff0019909116610100171660011790555b801562000493576000805461ff00191690555b50565b600054610100900460ff1680620004bb5750620004bb6001600160e01b03620001c516565b80620004ca575060005460ff16155b620005075760405162461bcd60e51b815260040180806020018281038252602e81526020018062002618602e913960400191505060405180910390fd5b600054610100900460ff1615801562000533576000805460ff1961ff0019909116610100171660011790555b82516020808501919091208351918401919091206065919091556066558015620003c2576000805461ff0019169055505050565b600054610100900460ff16806200058c57506200058c6001600160e01b03620001c516565b806200059b575060005460ff16155b620005d85760405162461bcd60e51b815260040180806020018281038252602e81526020018062002618602e913960400191505060405180910390fd5b600054610100900460ff1615801562000604576000805460ff1961ff0019909116610100171660011790555b604051806052620025c68239604051908190036052019020609a55508015620002e6576000805461ff00191690555050565b600054610100900460ff16806200065b57506200065b6001600160e01b03620001c516565b806200066a575060005460ff16155b620006a75760405162461bcd60e51b815260040180806020018281038252602e81526020018062002618602e913960400191505060405180910390fd5b600054610100900460ff16158015620006d3576000805460ff1961ff0019909116610100171660011790555b8251620006e890603690602086019062000723565b508151620006fe90603790602085019062000723565b506038805460ff191660121790558015620003c2576000805461ff0019169055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200076657805160ff191683800117855562000796565b8280016001018555821562000796579182015b828111156200079657825182559160200191906001019062000779565b50620007a4929150620007a8565b5090565b620001e191905b80821115620007a45760008155600101620007af565b611df180620007d56000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806374f4f547116100b8578063a708acce1161007c578063a708acce1461055c578063a9059cbb14610564578063c2eeeebd14610590578063d505accf146105b4578063dd62ed3e14610605578063e78cea921461063357610137565b806374f4f547146104aa5780637ecebe00146104d65780638c2a993e146104fc57806395d89b4114610528578063a457c2d71461053057610137565b8063313ce567116100ff578063313ce567146103795780633644e51514610397578063395093511461039f5780634000aea0146103cb57806370a082311461048457610137565b806306fdde031461013c578063095ea7b3146101b95780631624f6c6146101f957806318160ddd1461032957806323b872dd14610343575b600080fd5b61014461063b565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017e578181015183820152602001610166565b50505050905090810190601f1680156101ab5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e5600480360360408110156101cf57600080fd5b506001600160a01b0381351690602001356106d2565b604080519115158252519081900360200190f35b6103276004803603606081101561020f57600080fd5b810190602081018135600160201b81111561022957600080fd5b82018360208201111561023b57600080fd5b803590602001918460018302840111600160201b8311171561025c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156102ae57600080fd5b8201836020820111156102c057600080fd5b803590602001918460018302840111600160201b831117156102e157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff1691506106ef9050565b005b6103316107b0565b60408051918252519081900360200190f35b6101e56004803603606081101561035957600080fd5b506001600160a01b038135811691602081013590911690604001356107b6565b610381610843565b6040805160ff9092168252519081900360200190f35b61033161084c565b6101e5600480360360408110156103b557600080fd5b506001600160a01b03813516906020013561085b565b6101e5600480360360608110156103e157600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561041057600080fd5b82018360208201111561042257600080fd5b803590602001918460018302840111600160201b8311171561044357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108af945050505050565b6103316004803603602081101561049a57600080fd5b50356001600160a01b031661098a565b610327600480360360408110156104c057600080fd5b506001600160a01b0381351690602001356109a5565b610331600480360360208110156104ec57600080fd5b50356001600160a01b0316610a00565b6103276004803603604081101561051257600080fd5b506001600160a01b038135169060200135610a27565b610144610a7e565b6101e56004803603604081101561054657600080fd5b506001600160a01b038135169060200135610adf565b610327610b4d565b6101e56004803603604081101561057a57600080fd5b506001600160a01b038135169060200135610b4f565b610598610b63565b604080516001600160a01b039092168252519081900360200190f35b610327600480360360e08110156105ca57600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c00135610b72565b6103316004803603604081101561061b57600080fd5b506001600160a01b0381358116916020013516610d0a565b610598610d35565b60368054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106c75780601f1061069c576101008083540402835291602001916106c7565b820191906000526020600020905b8154815290600101906020018083116106aa57829003601f168201915b505050505090505b90565b60006106e66106df610d4a565b8484610d4e565b50600192915050565b600054610100900460ff16806107085750610708610e3a565b80610716575060005460ff16155b6107515760405162461bcd60e51b815260040180806020018281038252602e815260200180611c43602e913960400191505060405180910390fd5b600054610100900460ff1615801561077c576000805460ff1961ff0019909116610100171660011790555b61078584610e4b565b61078f8484610f21565b61079882610fd7565b80156107aa576000805461ff00191690555b50505050565b60355490565b60006107c3848484610fed565b610839846107cf610d4a565b61083485604051806060016040528060288152602001611ce5602891396001600160a01b038a1660009081526034602052604081209061080d610d4a565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61114416565b610d4e565b5060019392505050565b60385460ff1690565b60006108566111db565b905090565b60006106e6610868610d4a565b846108348560346000610879610d4a565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61120e16565b60006108bb8484610b4f565b50836001600160a01b0316336001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1685856040518083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561093657818101518382015260200161091e565b50505050905090810190601f1680156109635780820380516001836020036101000a031916815260200191505b50935050505060405180910390a361097a84610d44565b156108395761083984848461126f565b6001600160a01b031660009081526033602052604090205490565b60cc546001600160a01b031633146109f2576040805162461bcd60e51b815260206004820152600b60248201526a4f4e4c595f42524944474560a81b604482015290519081900360640190fd5b6109fc8282611349565b5050565b6001600160a01b0381166000908152609960205260408120610a219061143f565b92915050565b60cc546001600160a01b03163314610a74576040805162461bcd60e51b815260206004820152600b60248201526a4f4e4c595f42524944474560a81b604482015290519081900360640190fd5b6109fc8282611443565b60378054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106c75780601f1061069c576101008083540402835291602001916106c7565b60006106e6610aec610d4a565b8461083485604051806060016040528060258152602001611d976025913960346000610b16610d4a565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61114416565b565b60006106e6610b5c610d4a565b8484610fed565b60cd546001600160a01b031681565b83421115610bc7576040805162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e65000000604482015290519081900360640190fd5b6000609a54888888610bfc609960008e6001600160a01b03166001600160a01b0316815260200190815260200160002061143f565b604080516020808201979097526001600160a01b0395861681830152939094166060840152608083019190915260a082015260c08082018990528251808303909101815260e0909101909152805191012090506000610c5a8261152f565b90506000610c6a8287878761157b565b9050896001600160a01b0316816001600160a01b031614610cd2576040805162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e61747572650000604482015290519081900360640190fd5b6001600160a01b038a166000908152609960205260409020610cf3906116e6565b610cfe8a8a8a610d4e565b50505050505050505050565b6001600160a01b03918216600090815260346020908152604080832093909416825291909152205490565b60cc546001600160a01b031681565b3b151590565b3390565b6001600160a01b038316610d935760405162461bcd60e51b8152600401808060200182810382526024815260200180611d736024913960400191505060405180910390fd5b6001600160a01b038216610dd85760405162461bcd60e51b8152600401808060200182810382526022815260200180611b876022913960400191505060405180910390fd5b6001600160a01b03808416600081815260346020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6000610e4530610d44565b15905090565b600054610100900460ff1680610e645750610e64610e3a565b80610e72575060005460ff16155b610ead5760405162461bcd60e51b815260040180806020018281038252602e815260200180611c43602e913960400191505060405180910390fd5b600054610100900460ff16158015610ed8576000805460ff1961ff0019909116610100171660011790555b610ee06116ef565b610f0382604051806040016040528060018152602001603160f81b815250611791565b610f0c82611851565b80156109fc576000805461ff00191690555050565b600054610100900460ff1680610f3a5750610f3a610e3a565b80610f48575060005460ff16155b610f835760405162461bcd60e51b815260040180806020018281038252602e815260200180611c43602e913960400191505060405180910390fd5b600054610100900460ff16158015610fae576000805460ff1961ff0019909116610100171660011790555b610fb66116ef565b610fc0838361190e565b8015610fd2576000805461ff00191690555b505050565b6038805460ff191660ff92909216919091179055565b6001600160a01b0383166110325760405162461bcd60e51b8152600401808060200182810382526025815260200180611d4e6025913960400191505060405180910390fd5b6001600160a01b0382166110775760405162461bcd60e51b8152600401808060200182810382526023815260200180611b426023913960400191505060405180910390fd5b611082838383610fd2565b6110c581604051806060016040528060268152602001611ba9602691396001600160a01b038616600090815260336020526040902054919063ffffffff61114416565b6001600160a01b0380851660009081526033602052604080822093909355908416815220546110fa908263ffffffff61120e16565b6001600160a01b038084166000818152603360209081526040918290209490945580518581529051919392871692600080516020611d0d83398151915292918290030190a3505050565b600081848411156111d35760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611198578181015183820152602001611180565b50505050905090810190601f1680156111c55780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60006108566040518080611c9360529139605201905060405180910390206112016119e6565b6112096119ec565b6119f2565b600082820183811015611268576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b604051635260769b60e11b815233600482018181526024830185905260606044840190815284516064850152845187946001600160a01b0386169463a4c0ed369490938993899360840190602085019080838360005b838110156112dd5781810151838201526020016112c5565b50505050905090810190601f16801561130a5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561132b57600080fd5b505af115801561133f573d6000803e3d6000fd5b5050505050505050565b6001600160a01b03821661138e5760405162461bcd60e51b8152600401808060200182810382526021815260200180611d2d6021913960400191505060405180910390fd5b61139a82600083610fd2565b6113dd81604051806060016040528060228152602001611b65602291396001600160a01b038516600090815260336020526040902054919063ffffffff61114416565b6001600160a01b038316600090815260336020526040902055603554611409908263ffffffff611a4816565b6035556040805182815290516000916001600160a01b03851691600080516020611d0d8339815191529181900360200190a35050565b5490565b6001600160a01b03821661149e576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6114aa60008383610fd2565b6035546114bd908263ffffffff61120e16565b6035556001600160a01b0382166000908152603360205260409020546114e9908263ffffffff61120e16565b6001600160a01b0383166000818152603360209081526040808320949094558351858152935192939192600080516020611d0d8339815191529281900390910190a35050565b60006115396111db565b82604051602001808061190160f01b81525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050919050565b60006fa2a8918ca85bafe22016d0b997e4df60600160ff1b038211156115d25760405162461bcd60e51b8152600401808060200182810382526022815260200180611bcf6022913960400191505060405180910390fd5b8360ff16601b14806115e757508360ff16601c145b6116225760405162461bcd60e51b8152600401808060200182810382526022815260200180611c716022913960400191505060405180910390fd5b604080516000808252602080830180855289905260ff88168385015260608301879052608083018690529251909260019260a080820193601f1981019281900390910190855afa15801561167a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166116dd576040805162461bcd60e51b815260206004820152601860248201527745434453413a20696e76616c6964207369676e617475726560401b604482015290519081900360640190fd5b95945050505050565b80546001019055565b600054610100900460ff16806117085750611708610e3a565b80611716575060005460ff16155b6117515760405162461bcd60e51b815260040180806020018281038252602e815260200180611c43602e913960400191505060405180910390fd5b600054610100900460ff1615801561177c576000805460ff1961ff0019909116610100171660011790555b801561178e576000805461ff00191690555b50565b600054610100900460ff16806117aa57506117aa610e3a565b806117b8575060005460ff16155b6117f35760405162461bcd60e51b815260040180806020018281038252602e815260200180611c43602e913960400191505060405180910390fd5b600054610100900460ff1615801561181e576000805460ff1961ff0019909116610100171660011790555b82516020808501919091208351918401919091206065919091556066558015610fd2576000805461ff0019169055505050565b600054610100900460ff168061186a575061186a610e3a565b80611878575060005460ff16155b6118b35760405162461bcd60e51b815260040180806020018281038252602e815260200180611c43602e913960400191505060405180910390fd5b600054610100900460ff161580156118de576000805460ff1961ff0019909116610100171660011790555b604051806052611bf18239604051908190036052019020609a555080156109fc576000805461ff00191690555050565b600054610100900460ff16806119275750611927610e3a565b80611935575060005460ff16155b6119705760405162461bcd60e51b815260040180806020018281038252602e815260200180611c43602e913960400191505060405180910390fd5b600054610100900460ff1615801561199b576000805460ff1961ff0019909116610100171660011790555b82516119ae906036906020860190611aa9565b5081516119c2906037906020850190611aa9565b506038805460ff191660121790558015610fd2576000805461ff0019169055505050565b60655490565b60665490565b60008383836119ff611aa5565b6040805160208082019690965280820194909452606084019290925260808301523060a0808401919091528151808403909101815260c090920190528051910120949350505050565b600082821115611a9f576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b4690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611aea57805160ff1916838001178555611b17565b82800160010185558215611b17579182015b82811115611b17578251825591602001919060010190611afc565b50611b23929150611b27565b5090565b6106cf91905b80821115611b235760008155600101611b2d56fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545434453413a20696e76616c6964207369676e6174757265202773272076616c75655065726d69742861646472657373206f776e65722c61646472657373207370656e6465722c75696e743235362076616c75652c75696e74323536206e6f6e63652c75696e7432353620646561646c696e6529496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a656445434453413a20696e76616c6964207369676e6174757265202776272076616c7565454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e74726163742945524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef45524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220b834ac18a7fe2da6a606ecd6088ec95f60f23edd775a703289999b9ea476252464736f6c634300060b00335065726d69742861646472657373206f776e65722c61646472657373207370656e6465722c75696e743235362076616c75652c75696e74323536206e6f6e63652c75696e7432353620646561646c696e6529496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564'
