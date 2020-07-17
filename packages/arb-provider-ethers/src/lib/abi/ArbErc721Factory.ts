/* Generated by ts-generator ver. 0.0.8 */
/* tslint:disable */

import { Contract, ContractFactory, Signer } from 'ethers'
import { Provider } from 'ethers/providers'
import { UnsignedTransaction } from 'ethers/utils/transaction'

import { TransactionOverrides } from '.'
import { ArbErc721 } from './ArbErc721'

export class ArbErc721Factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(overrides?: TransactionOverrides): Promise<ArbErc721> {
    return super.deploy(overrides) as Promise<ArbErc721>
  }
  getDeployTransaction(overrides?: TransactionOverrides): UnsignedTransaction {
    return super.getDeployTransaction(overrides)
  }
  attach(address: string): ArbErc721 {
    return super.attach(address) as ArbErc721
  }
  connect(signer: Signer): ArbErc721Factory {
    return super.connect(signer) as ArbErc721Factory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ArbErc721 {
    return new Contract(address, _abi, signerOrProvider) as ArbErc721
  }
}

const _abi = [
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
        name: 'approved',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'uint256',
        name: 'tokenId',
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
        name: 'owner',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'operator',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'bool',
        name: 'approved',
        type: 'bool',
      },
    ],
    name: 'ApprovalForAll',
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
        indexed: true,
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
    ],
    name: 'Transfer',
    type: 'event',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: 'account',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
    ],
    name: 'adminMint',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
    ],
    name: 'approve',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: 'owner',
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
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
    ],
    name: 'getApproved',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: 'owner',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'operator',
        type: 'address',
      },
    ],
    name: 'isApprovedForAll',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
    ],
    name: 'ownerOf',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
    ],
    name: 'safeTransferFrom',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_data',
        type: 'bytes',
      },
    ],
    name: 'safeTransferFrom',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        internalType: 'bool',
        name: 'approved',
        type: 'bool',
      },
    ],
    name: 'setApprovalForAll',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'bytes4',
        name: 'interfaceId',
        type: 'bytes4',
      },
    ],
    name: 'supportsInterface',
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'uint256',
        name: 'index',
        type: 'uint256',
      },
    ],
    name: 'tokenByIndex',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: 'owner',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'index',
        type: 'uint256',
      },
    ],
    name: 'tokenOfOwnerByIndex',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [
      {
        internalType: 'address',
        name: 'owner',
        type: 'address',
      },
    ],
    name: 'tokensOfOwner',
    outputs: [
      {
        internalType: 'uint256[]',
        name: '',
        type: 'uint256[]',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: true,
    inputs: [],
    name: 'totalSupply',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: 'from',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
    ],
    name: 'transferFrom',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    constant: false,
    inputs: [
      {
        internalType: 'address',
        name: 'account',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
    ],
    name: 'withdraw',
    outputs: [],
    payable: false,
    stateMutability: 'nonpayable',
    type: 'function',
  },
]

const _bytecode =
  '0x608060405261001d6301ffc9a760e01b6001600160e01b0361005416565b6100366380ac58cd60e01b6001600160e01b0361005416565b61004f63780e9d6360e01b6001600160e01b0361005416565b6100d8565b6001600160e01b031980821614156100b3576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b611848806100e76000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80636352211e11610097578063b88d4fde11610066578063b88d4fde1461035d578063e58306f914610423578063e985e9c51461044f578063f3fef3a31461047d57610100565b80636352211e1461027657806370a08231146102935780638462151c146102b9578063a22cb4651461032f57610100565b806323b872dd116100d357806323b872dd146101c15780632f745c59146101f757806342842e0e146102235780634f6ccce71461025957610100565b806301ffc9a714610105578063081812fc14610140578063095ea7b31461017957806318160ddd146101a7575b600080fd5b61012c6004803603602081101561011b57600080fd5b50356001600160e01b0319166104a9565b604080519115158252519081900360200190f35b61015d6004803603602081101561015657600080fd5b50356104c8565b604080516001600160a01b039092168252519081900360200190f35b6101a56004803603604081101561018f57600080fd5b506001600160a01b03813516906020013561052a565b005b6101af610652565b60408051918252519081900360200190f35b6101a5600480360360608110156101d757600080fd5b506001600160a01b03813581169160208101359091169060400135610659565b6101af6004803603604081101561020d57600080fd5b506001600160a01b0381351690602001356106b5565b6101a56004803603606081101561023957600080fd5b506001600160a01b03813581169160208101359091169060400135610734565b6101af6004803603602081101561026f57600080fd5b503561074f565b61015d6004803603602081101561028c57600080fd5b50356107b5565b6101af600480360360208110156102a957600080fd5b50356001600160a01b031661080f565b6102df600480360360208110156102cf57600080fd5b50356001600160a01b0316610877565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561031b578181015183820152602001610303565b505050509050019250505060405180910390f35b6101a56004803603604081101561034557600080fd5b506001600160a01b03813516906020013515156108d8565b6101a56004803603608081101561037357600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156103ae57600080fd5b8201836020820111156103c057600080fd5b803590602001918460018302840111640100000000831117156103e257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109dd945050505050565b6101a56004803603604081101561043957600080fd5b506001600160a01b038135169060200135610a3b565b61012c6004803603604081101561046557600080fd5b506001600160a01b0381358116916020013516610a56565b6101a56004803603604081101561049357600080fd5b506001600160a01b038135169060200135610a84565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006104d382610af9565b61050e5760405162461bcd60e51b815260040180806020018281038252602c81526020018061171c602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b6000610535826107b5565b9050806001600160a01b0316836001600160a01b031614156105885760405162461bcd60e51b81526004018080602001828103825260218152602001806117716021913960400191505060405180910390fd5b806001600160a01b031661059a610b16565b6001600160a01b031614806105bb57506105bb816105b6610b16565b610a56565b6105f65760405162461bcd60e51b81526004018080602001828103825260388152602001806116916038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6007545b90565b61066a610664610b16565b82610b1a565b6106a55760405162461bcd60e51b81526004018080602001828103825260318152602001806117926031913960400191505060405180910390fd5b6106b0838383610bbe565b505050565b60006106c08361080f565b82106106fd5760405162461bcd60e51b815260040180806020018281038252602b8152602001806115e4602b913960400191505060405180910390fd5b6001600160a01b038316600090815260056020526040902080548390811061072157fe5b9060005260206000200154905092915050565b6106b0838383604051806020016040528060008152506109dd565b6000610759610652565b82106107965760405162461bcd60e51b815260040180806020018281038252602c8152602001806117c3602c913960400191505060405180910390fd5b600782815481106107a357fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b0316806108095760405162461bcd60e51b81526004018080602001828103825260298152602001806116f36029913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166108565760405162461bcd60e51b815260040180806020018281038252602a8152602001806116c9602a913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061080990610bdd565b606061088282610be1565b8054806020026020016040519081016040528092919081815260200182805480156108cc57602002820191906000526020600020905b8154815260200190600101908083116108b8575b50505050509050919050565b6108e0610b16565b6001600160a01b0316826001600160a01b03161415610946576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b8060046000610953610b16565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff191692151592909217909155610997610b16565b60408051841515815290516001600160a01b0392909216917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319181900360200190a35050565b6109ee6109e8610b16565b83610b1a565b610a295760405162461bcd60e51b81526004018080602001828103825260318152602001806117926031913960400191505060405180910390fd5b610a3584848484610bfb565b50505050565b33600114610a4857600080fd5b610a528282610c4d565b5050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b610a8e3382610c6a565b60408051631e7c829f60e31b81526001600160a01b038416600482015260248101839052905160649163f3e414f891604480830192600092919082900301818387803b158015610add57600080fd5b505af1158015610af1573d6000803e3d6000fd5b505050505050565b6000908152600160205260409020546001600160a01b0316151590565b3390565b6000610b2582610af9565b610b605760405162461bcd60e51b815260040180806020018281038252602c815260200180611665602c913960400191505060405180910390fd5b6000610b6b836107b5565b9050806001600160a01b0316846001600160a01b03161480610ba65750836001600160a01b0316610b9b846104c8565b6001600160a01b0316145b80610bb65750610bb68185610a56565b949350505050565b610bc9838383610c96565b610bd38382610dda565b6106b08282610ecf565b5490565b6001600160a01b0316600090815260056020526040902090565b610c06848484610bbe565b610c1284848484610f0d565b610a355760405162461bcd60e51b815260040180806020018281038252603281526020018061160f6032913960400191505060405180910390fd5b610c578282611148565b610c618282610ecf565b610a5281611279565b610c7482826112bd565b610c7e8282610dda565b600081815260066020526040812055610a5281611394565b826001600160a01b0316610ca9826107b5565b6001600160a01b031614610cee5760405162461bcd60e51b81526004018080602001828103825260298152602001806117486029913960400191505060405180910390fd5b6001600160a01b038216610d335760405162461bcd60e51b81526004018080602001828103825260248152602001806116416024913960400191505060405180910390fd5b610d3c81611430565b6001600160a01b0383166000908152600360205260409020610d5d9061146d565b6001600160a01b0382166000908152600360205260409020610d7e90611484565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054610e0490600163ffffffff61148d16565b600083815260066020526040902054909150808214610e9f576001600160a01b0384166000908152600560205260408120805484908110610e4157fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b031681526020019081526020016000208381548110610e7f57fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490610ec89060001983016115a6565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b6000610f21846001600160a01b03166114d6565b610f2d57506001610bb6565b600060606001600160a01b038616630a85bd0160e11b610f4b610b16565b89888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610fc4578181015183820152602001610fac565b50505050905090810190601f168015610ff15780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b602083106110595780518252601f19909201916020918201910161103a565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146110bb576040519150601f19603f3d011682016040523d82523d6000602084013e6110c0565b606091505b509150915081611111578051156110da5780518082602001fd5b60405162461bcd60e51b815260040180806020018281038252603281526020018061160f6032913960400191505060405180910390fd5b600081806020019051602081101561112857600080fd5b50516001600160e01b031916630a85bd0160e11b149350610bb692505050565b6001600160a01b0382166111a3576040805162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015290519081900360640190fd5b6111ac81610af9565b156111fe576040805162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015290519081900360640190fd5b600081815260016020908152604080832080546001600160a01b0319166001600160a01b03871690811790915583526003909152902061123d90611484565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b600780546000838152600860205260408120829055600182018355919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880155565b816001600160a01b03166112d0826107b5565b6001600160a01b0316146113155760405162461bcd60e51b81526004018080602001828103825260258152602001806117ef6025913960400191505060405180910390fd5b61131e81611430565b6001600160a01b038216600090815260036020526040902061133f9061146d565b60008181526001602052604080822080546001600160a01b0319169055518291906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6007546000906113ab90600163ffffffff61148d16565b600083815260086020526040812054600780549394509092849081106113cd57fe5b9060005260206000200154905080600783815481106113e857fe5b6000918252602080832090910192909255828152600890915260409020829055600780549061141b9060001983016115a6565b50505060009182525060086020526040812055565b6000818152600260205260409020546001600160a01b03161561146a57600081815260026020526040902080546001600160a01b03191690555b50565b805461148090600163ffffffff61148d16565b9055565b80546001019055565b60006114cf83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061150f565b9392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610bb6575050151592915050565b6000818484111561159e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561156357818101518382015260200161154b565b50505050905090810190601f1680156115905780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b8154818355818111156106b0576000838152602090206106b091810190830161065691905b808211156115df57600081556001016115cb565b509056fe455243373231456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564455243373231456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e64734552433732313a206275726e206f6620746f6b656e2074686174206973206e6f74206f776ea265627a7a723158208d1692188bcbe9e434c19268ca7dba3f00c775ef324fe1ed64d899530870fa4264736f6c63430005110032'
