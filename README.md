<p align="center">
  <a href="https://www.riotgames.com"><img src="https://www.riotgames.com/darkroom/800/4055f068b0bc39bbe60d08491a39994b:0ff2ceae3e7bed5e6f2c9c5a1aded2c0/riot-fist-container-white-red.png" width="250" title="Riot Fist Container White (Red)"></a>
</p>

# Go Riot API

## 🎯 Table of Contents

*	[Installation](#installation)
* [Getting started](#getting-started)
* [Usage](#usage)
  *	[GetAccountByRiotId](#getAccountByRiotId)
  *	[GetAccountByPuuid](#getAccountByPuuid)
  *	[GetActiveShard](#getActiveShard)

<a name="installation"></a>

## ⚙️ Installation

```
Get:
go get github.com/guillaumeguirriec/go-riot-api

Import:
import("github.com/guillaumeguirriec/go-riot-api")
```

<a name="getting-started"></a>

## ⚡️ Getting started

<a name="usage"></a>

## 👀 Usage

<a name="getAccountByRiotId"></a>

### 📖 GetAccountByRiotId

```go
response, err := riot.GetAccountByRiotId("smile", "6578")

if err != nil {
	fmt.Println(err)
} else {
	PrintResult(response)
}
```

<details>
  <summary>JSON response</summary>

```go
{
    "puuid": "jUr1OkZKAS4AW6amrpxOmfYin3w9P-jiVuI7UtNmyrJRL9Z5B0R_Qzs6h7pEwCThABtBODsoyhcDbQ",
    "gameName": "smile",
    "tagLine": "6578"
}
```
</details>

<details>
  <summary>Struct definition</summary>

```go
type AccountDto struct {
  Puuid, GameName, TagLine string
}
```
</details>

<a name="getAccountByPuuid"></a>

### 📖 GetAccountByPuuid

```go
response, err := riot.GetAccountByPuuid("jUr1OkZKAS4AW6amrpxOmfYin3w9P-jiVuI7UtNmyrJRL9Z5B0R_Qzs6h7pEwCThABtBODsoyhcDbQ")

if err != nil {
	fmt.Println(err)
} else {
	PrintResult(response)
}
```

<details>
  <summary>JSON response</summary>

```go
{
    "puuid": "jUr1OkZKAS4AW6amrpxOmfYin3w9P-jiVuI7UtNmyrJRL9Z5B0R_Qzs6h7pEwCThABtBODsoyhcDbQ",
    "gameName": "smile",
    "tagLine": "6578"
}
```
</details>

<details>
  <summary>Struct definition</summary>

```go
type AccountDto struct {
  Puuid, GameName, TagLine string
}
```
</details>

<a name="getActiveShard"></a>

### 📖 GetActiveShard

```go
response, err := riot.GetActiveShard("jUr1OkZKAS4AW6amrpxOmfYin3w9P-jiVuI7UtNmyrJRL9Z5B0R_Qzs6h7pEwCThABtBODsoyhcDbQ", "val")

if err != nil {
	fmt.Println(err)
} else {
	PrintResult(response)
}
```

<details>
  <summary>JSON response</summary>

```go
{
    "puuid": "jUr1OkZKAS4AW6amrpxOmfYin3w9P-jiVuI7UtNmyrJRL9Z5B0R_Qzs6h7pEwCThABtBODsoyhcDbQ",
    "game": "val",
    "activeShard": "eu"
}
```
</details>

<details>
  <summary>Struct definition</summary>

```go
type ActiveShardDto struct {
	Puuid, Game, ActiveShard string
}
```
</details>
