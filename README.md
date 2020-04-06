<div align="center">
  
![GitHub contributors](https://img.shields.io/github/contributors/JosephZoeller/maritime-royale)
![GitHub forks](https://img.shields.io/github/forks/JosephZoeller/maritime-royale?label=Forks)
![GitHub stars](https://img.shields.io/github/stars/JosephZoeller/maritime-royale?style=Stars)
![GitHub issues](https://img.shields.io/github/issues-raw/JosephZoeller/maritime-royale)
[![Go Report Card](https://goreportcard.com/badge/github.com/JosephZoeller/maritime-royale)](https://goreportcard.com/report/github.com/JosephZoeller/maritime-royale)

</div>

<h3 align="center">
    
  ![destroyer_img](https://cdn.discordapp.com/attachments/689340216284020812/689360748920832000/destroyer.png) Maritime Royale ![destroyer_img](https://cdn.discordapp.com/attachments/689340216284020812/689360748920832000/destroyer.png)
  
A battle-royale-style, synchronous turn-based strategy game.
</h3>


<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Project Proposal](#project-proposal)
  * [Game Overview](#game-overview)
  * [Build from Source](#installation)
  * [Built With](#built-with)
  
<!-- ABOUT THE PROJECT -->
# About The Project 

## Project Proposal
An ambitious Player vs. Player strategy game set on a dynamically sized 2D grid, where player-controlled units contend for scarce unit-construction resources in an effort to gain total control of the game map.

## Game Overview
At the beginning of the game, all players are allotted with a Mothership: a slow-moving capital unit which can attack enemy-controlled units and construct player-controlled Airship units. When a player's Mothership is destroyed, the player is defeated. Thus, the player's primary objective is to destroy all enemy Motherships.

The game map is a 2-Dimensional grid, dynamically sized to fit the number of participants. Each player Mothership is randomly placed on a valid tile (non-occupied, water tile). Once all Motherships have been placed, the game is begun.

Each round is time-limited and carried out asynchronously for all players, meaning that the players are performing their turns at the same time. For each turn, every player-controlled unit has the opportunity to move and perform a single action. These actions include:
* Wait - The unit can forego their action to await the next round.
* Attack - The unit can engage an enemy unit in combat.
* Capture - Certain units can lay claim to special map "resource" tiles.

Once an action is performed by a unit, the player must wait until the next round to control that unit again. Player-controllable units include:
* Mothership - A slow-moving aircraft carrier that can construct Jets and capture resource tiles
* Jets - An aircraft unit that can easily defeat submarines, but have a difficult time against Destroyers.
* Destroyers - A naval unit which can combat Jets, but are susceptible to Submarine attacks.
* Submarines - A subnautical unit designed to eliminate Destroyers, but are thwarted by Jets.

The game map is predominantly constituted of water tiles, where units can move freely and engage in naval battle. However, prior to the Mothership placement phase, special resource tiles are sparsely distributed across the map. Players can claim and reclaim these tiles by using specific unit-types to capture them. Once claimed, these tiles will grant benefits to the owner. These tiles, and their benefits, include:
* Islands - Each turn, awards the owner with currency. Also increases the max capacity of the owner's fleet.
* Shipyards - Each turn, the owner can spend currency to construct a ship on this tile. The unit is added to the owner's fleet.

Players are tasked with strategically combatting enemy players while capturing, managing and protecting map resources. The last surviving player wins the match!

# Technology

## Overview

The MARI Engine on demonstration here extends the [Ebiten](https://ebiten.org/) windowing/game engine to use data oriented programming or DOP. It does this through the use of the Element and Component. The MARI Engine also can replicate any and all behavior from clients through its multiplayer framework. This scaling multiplayer network can handle large traffic volumes while maintaining deliverability through smart updating.

### DOP

The main idea of DOP design is that we seperate data with the functions that call that data in the most generic way. The purpose of doing this is to prevent upcalling parent classes for data. This approach a more  efficent operation for the CPU to perform. Alongside the this improvement, the DOP design structure creates extremly generic code. It is so generic in fact that DOP code can be all loaded into a single structure. 

### Elements

The Element is a structure found inside the MARI engine framework. Every object in the MARI engine is an Element. The Element is composed of a few default (almost universally used) data types. 

*     XPos       float64
Stores the world x position of the top left cornor of the Element
*     YPos       float64
Stores the world y position of the top left cornor of the Element
*     Rotation   float64
Stores the world rotation of the Element
*     Active     bool
Determines if the Element should be enabled or not. This is not for image culling but rather for setting an Element to be completly dorment. 
*     UniqueName string
The unique identifier for the Element. No two Elements may share a UniqueName during an online session. This is a nil value for all single player/non server related launches.
*     ID         string
Determines the owner of the Element. 
*     Components []Component
The extention of Element. The component slice is what adds all functionality to the Element.

### Components

    type Component interface {
      OnUpdate(world []*Element) error
      OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error
      OnCheck(*Element) error
      OnUpdateServer(world []*Element) error
      MRP(finalElem *Element, conn net.Conn)
    }

A component is simply anything fitting the abover interface. At first this may not seem impactful, but by having all objects in a game be the same structure allows for massive flexability. 

Explanation by example is preferred. Pretend you have lots of trees and bushes in your game. Suddenly it comes upon you that your game must have fire physics. That fire should be able to spread from tree to tree and from tree to bush and bush to bush.

Before to accomplish this, you would have to write functions for each object you wanted to immplemnt fire physics for and then assign each object its variables. This is tiresome escpecially if the physics needs to change. This would require changing every place you implemented code relating to fire physics. There are some ways around this dilema with OOP design, but none are very elegant, nor do they scale very well.

Here is where DOP shines. We create one component "FirePhysics" and write the code it needs to call for in the update or other respective functions. If we need variables custom to fire physics such as burnability and fire width, we can assign them to the custom component. Now we can add this component to anything and it will immediately implement our code. This kind of scalability is unparalleled. Not only this, but our code is now all in one location allowing for easy source control. 

#### Kinds of Components

There are infinitly many kinds of components, but there are only three main ones

* First Order
* Second Order
* Third Order

The naming convention beyond this is self explanatory. 

## Decalaring New Components 

## Installation

If you are adding assets to the game, please make sure they are statically compiled with the binary for portability. Follow these instruction on how to
do so:

First install the necessary "statik" tool to convert assets.

```$ go get github.com/rakyll/statik```

Then run this command before ```go build``` of the server or client while in the root directory of the project workspace.

```$ statik -src=./assets/sprites```

Now you can perform the compilation.

```$ go build ./cmd/server/```
```$ go build ./cmd/client/```

### Built With

* [Go](https://golang.org/)
* [AWS](https://aws.amazon.com/)
* [Ubuntu 18.04](http://releases.ubuntu.com/18.04.4/)
* [Docker](https://www.docker.com/)
* [Kubernetes](https://kubernetes.io/)
* [Terraform](https://www.terraform.io/)
* [go-sdl2](https://github.com/veandco/go-sdl2)


![Demo Image](https://cdn.discordapp.com/attachments/689340216284020812/689544753075060736/screenshot.png)
