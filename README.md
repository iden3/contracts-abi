# Contract's abi
This repository contains contract interfaces for various languages.

**!!Important!!**

State contract abi is stable, and in minor versions new functionality can be added.
Meantime Onchain Resovler and Identity are not stable yet and ABI can be updated in minor versions of package. 

## Requirements:
```
jq
npm
```

Before running the `./generator` script, you need to clone the [contracts repository](https://github.com/iden3/contracts) or create a symlink to it:
```bash
ln -s $PATH_TO_CONTRACTS contracts
```

## To run the script, follow these steps:
```bash
./generator
```

## Method ids

### Onchain non merklized issuer base

| Method signature                     | Method id  |
|--------------------------------------|------------|
| INonMerklizedIssuer                  | 0x58874949 |
| supportsInterface(bytes4)            | 0x01ffc9a7 |
| getUserCredentialIds(uint256)        | 0x668d0bd4 |
| getCredential(uint256,uint256)       | 0x37c1d9ff |
| getCredentialAdapterVersion()        | 0x09cb9b62 |

### IZKPVerifier

| Method signature                     | Method id  |
|--------------------------------------|------------|
| getZKPRequestsCount()                | 0x6508e1b4 |
| getZKPRequest(uint64)                | 0xc76d0845 |
| getZKPRequests(uint256,uint256)      | 0x5f9e60d7 |

## Contributing

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
dual licensed as below, without any additional terms or conditions.

## License

Copyright 2023 0kims Association

This project is licensed under either of

- [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE))
- [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT))

at your option.
