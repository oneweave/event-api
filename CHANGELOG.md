# Changelog

## [1.18.0](https://github.com/oneweave/event-api/compare/v1.17.0...v1.18.0) (2026-06-22)


### Features

* add method to create envelope from an existing envelope ([e6fe541](https://github.com/oneweave/event-api/commit/e6fe54139665a94512cdc33a3f1ebfdede595c68))
* update cloud event building ([e386983](https://github.com/oneweave/event-api/commit/e386983718831c76e5ac1d75f6dfd529e36cb7cf))

## [1.17.0](https://github.com/oneweave/event-api/compare/v1.16.0...v1.17.0) (2026-06-22)


### Features

* update controller update events ([3a179f1](https://github.com/oneweave/event-api/commit/3a179f1215d5b514f78378d35053c5574bf478d7))

## [1.16.0](https://github.com/oneweave/event-api/compare/v1.15.0...v1.16.0) (2026-06-22)


### Features

* add cloud event builder ([55c89ed](https://github.com/oneweave/event-api/commit/55c89ed9b3313254727131ef3575732450589e58))

## [1.15.0](https://github.com/oneweave/event-api/compare/v1.14.0...v1.15.0) (2026-06-21)


### Features

* add tenant status enum ([d196913](https://github.com/oneweave/event-api/commit/d196913b735cbf34260c42bf6cb29a27e1a05e67))

## [1.14.0](https://github.com/oneweave/event-api/compare/v1.13.0...v1.14.0) (2026-06-21)


### Features

* add tenant events and data ([fd4b25c](https://github.com/oneweave/event-api/commit/fd4b25ce6c1728ede1158fe6eee4f5ba146b3aee))

## [1.13.0](https://github.com/oneweave/event-api/compare/v1.12.0...v1.13.0) (2026-06-21)


### Features

* remove artifact id, not needed ([799df78](https://github.com/oneweave/event-api/commit/799df786168f6fe1d84dfe9a13f69d646d471c34))
* remove duplicated fields from build event data ([d138f60](https://github.com/oneweave/event-api/commit/d138f60813d7984fcd527044ef9f74a2605a9a9d))

## [1.12.0](https://github.com/oneweave/event-api/compare/v1.11.1...v1.12.0) (2026-06-21)


### Features

* add fn to parse and validate ([a9fdacf](https://github.com/oneweave/event-api/commit/a9fdacf2d2fe204bc9dd8fe77d138b26fb865eee))
* use our event-id instead of uuids ([7616ca3](https://github.com/oneweave/event-api/commit/7616ca3a0d69e585d907b97095f273a507f6175f))

## [1.11.1](https://github.com/oneweave/event-api/compare/v1.11.0...v1.11.1) (2026-06-21)


### Bug Fixes

* add missing release fields to build succeeded ([5e3def4](https://github.com/oneweave/event-api/commit/5e3def47b13e3bc5ad88a4a7fbbf3580ed8f102c))

## [1.11.0](https://github.com/oneweave/event-api/compare/v1.10.0...v1.11.0) (2026-06-21)


### Features

* add artifact build request ([053c3dd](https://github.com/oneweave/event-api/commit/053c3dd2b86d506cfe6e94963d657bd79a87e5c9))
* update asyncapi spec ([c0c003e](https://github.com/oneweave/event-api/commit/c0c003e3c8d742bd10a5aabbd04b06ea58daf613))
* update event payloads ([3d7bed4](https://github.com/oneweave/event-api/commit/3d7bed4502c8ffeec17c91b341bb42a5ec7b0bbb))

## [1.10.0](https://github.com/oneweave/event-api/compare/v1.9.0...v1.10.0) (2026-06-16)


### Features

* set default for manifest file ([984c74c](https://github.com/oneweave/event-api/commit/984c74cb9431fb3cafca020f596c02f2792fe66f))


### Bug Fixes

* remove broken validators ([4ed9a41](https://github.com/oneweave/event-api/commit/4ed9a41dde7d898b57ad6118db978d07366ee5cd))

## [1.9.0](https://github.com/oneweave/event-api/compare/v1.8.0...v1.9.0) (2026-06-16)


### Features

* improve validation ([652c63f](https://github.com/oneweave/event-api/commit/652c63fb693b45a18989553a1fa88544c3fcfaae))

## [1.8.0](https://github.com/oneweave/event-api/compare/v1.7.0...v1.8.0) (2026-06-16)


### Features

* add min and max replica configuration ([872bea5](https://github.com/oneweave/event-api/commit/872bea5ce813e1b2d0aa4b916b48bd96c214c1dc))


### Bug Fixes

* remove dive tags from structs ([1f23dfb](https://github.com/oneweave/event-api/commit/1f23dfbac055ff66a7002241225064fa387ce478))

## [1.7.0](https://github.com/oneweave/event-api/compare/v1.6.0...v1.7.0) (2026-06-14)


### Features

* update broker event names and data ([8bcf1fb](https://github.com/oneweave/event-api/commit/8bcf1fbe44dd97224c67b7bcc274aff7d871d56e))

## [1.6.0](https://github.com/oneweave/event-api/compare/v1.5.0...v1.6.0) (2026-06-14)


### Features

* add broker and controller events ([9566cda](https://github.com/oneweave/event-api/commit/9566cdaa205d1fb34889db425e371bee1c1350ef))

## [1.5.0](https://github.com/oneweave/event-api/compare/v1.4.1...v1.5.0) (2026-06-13)


### Features

* add build info to build events ([0cd1cbe](https://github.com/oneweave/event-api/commit/0cd1cbe7d1eb1644576040c2ef31d7d14b27258a))
* add release events and data ([66a1fa4](https://github.com/oneweave/event-api/commit/66a1fa46d3f4eae8b221619f580ec7f99b59e01c))
* remove source revision fields, add release id ([fe73b5a](https://github.com/oneweave/event-api/commit/fe73b5aee1b17df45d62a6b0a43d201036a54bd3))

## [1.4.1](https://github.com/oneweave/event-api/compare/v1.4.0...v1.4.1) (2026-06-08)


### Bug Fixes

* ensure we return a pull target array ([50a3a58](https://github.com/oneweave/event-api/commit/50a3a58d061e963933af6244a4a5f7f5347c5555))

## [1.4.0](https://github.com/oneweave/event-api/compare/v1.3.0...v1.4.0) (2026-06-08)


### Features

* add helper for pull target from push target creation ([495de05](https://github.com/oneweave/event-api/commit/495de05d291570c4dbe3a08236399a0498934aae))

## [1.3.0](https://github.com/oneweave/event-api/compare/v1.2.0...v1.3.0) (2026-06-08)


### Features

* ensure we use cloudevent v2 sdk apis ([f0842a8](https://github.com/oneweave/event-api/commit/f0842a8a2d316c5a49ad6e35eccd2ee16c1caf49))

## [1.2.0](https://github.com/oneweave/event-api/compare/v1.1.0...v1.2.0) (2026-06-08)


### Features

* add function to create envelope from cloud event ([58007ca](https://github.com/oneweave/event-api/commit/58007ca9bf200eb9697f07618d23284c823914e1))

## [1.1.0](https://github.com/oneweave/event-api/compare/v1.0.0...v1.1.0) (2026-06-08)


### Features

* rename method ([0c540c5](https://github.com/oneweave/event-api/commit/0c540c5ecbd6be71dc79fbbbf768208a83ec2b06))

## 1.0.0 (2026-06-08)


### Features

* add async api components ([677923c](https://github.com/oneweave/event-api/commit/677923c8ef6c75fa0d8f795c3f471932d5c0d099))
