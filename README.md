# Contract's abi

This repository contains contract interfaces for various languages.

**!!Important!!**

State contract abi is stable, and in minor versions new functionality can be added.
Meantime Onchain Resovler and Identity are not stable yet and ABI can be updated in minor versions of package.

## Requirements

```
jq
npm
```

Before running the `./generator` script, you need to clone the [contracts repository](https://github.com/iden3/contracts) or create a symlink to it:

```bash
ln -s $PATH_TO_CONTRACTS contracts
```

## To run the script for go generator, follow these steps

```bash
./generator
```
## Generation types for TypeScript
1. Copy `package.json`, `tsconfig.json` and `.gitignore` from existing project with `js` folder.
2. Change corresponding name in the `package.json`.
3. Install dependencies with `npm i`.
4. Create `src` folder.
5. Copy abi `.json` in `src` folder from compiled artifacts in `contracts/artifacts` from `iden3/contracts` hardhat project.
6. Run command `npx typechain --target=ethers-v6 src/<your_abi>.json --out-dir ./src` replacing `<your_abi>` with the correct filename.

## Method ids

### Onchain non merklized issuer base

| Method signature                     | Method id  |
|--------------------------------------|------------|
| IIdentifiable                        | 0x5d1ca631 |
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

