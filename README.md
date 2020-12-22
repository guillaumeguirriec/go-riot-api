<p align="center">
  <a href="https://www.riotgames.com"><img src="https://www.riotgames.com/darkroom/800/4055f068b0bc39bbe60d08491a39994b:0ff2ceae3e7bed5e6f2c9c5a1aded2c0/riot-fist-container-white-red.png" width="250" title="Riot Fist Container White (Red)"></a>
</p>

# Go Riot API

This is the Go wrapper for the ACCOUNT-V1 Riot API. For more information consult the [ACCOUNT-V1 Riot API documentation](https://developer.riotgames.com/apis#account-v1).

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

```go
package main

import (
	"fmt"
	"github.com/guillaumeguirriec/go-riot-api"
)

func main() {
	riot, err := riot.New("<API key>", riot.EuropeRegion, riot.HeaderParam, false)

	if err != nil {
		fmt.Println(err)
  	}
  
  ...
}
```

<a name="usage"></a>

## 👀 Usage

<a name="getAccountByRiotId"></a>

### 📖 GetAccountByRiotId

```go
response, err := riot.GetAccountByRiotId("smile", "6578")

if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
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

For detailed information about the endpoint called, please refer to the [ACCOUNT-V1 Riot API "Get account by riot id" method documentation](https://developer.riotgames.com/apis#account-v1/GET_getByRiotId).

<a name="getAccountByPuuid"></a>

### 📖 GetAccountByPuuid

```go
response, err := riot.GetAccountByPuuid("jUr1OkZKAS4AW6amrpxOmfYin3w9P-jiVuI7UtNmyrJRL9Z5B0R_Qzs6h7pEwCThABtBODsoyhcDbQ")

if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
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

For detailed information about the endpoint called, please refer to the [ACCOUNT-V1 Riot API "Get account by puuid" method documentation](https://developer.riotgames.com/apis#account-v1/GET_getByPuuid).

<a name="getActiveShard"></a>

### 📖 GetActiveShard

```go
response, err := riot.GetActiveShard("jUr1OkZKAS4AW6amrpxOmfYin3w9P-jiVuI7UtNmyrJRL9Z5B0R_Qzs6h7pEwCThABtBODsoyhcDbQ", "val")

if err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%+v\n", response)
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

For detailed information about the endpoint called, please refer to the [ACCOUNT-V1 Riot API "Get active shard for a player" method documentation](https://developer.riotgames.com/apis#account-v1/GET_getActiveShard).
