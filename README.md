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
  * [Built With](#built-with)
  
<!-- ABOUT THE PROJECT -->
## About The Project 

### Project Proposal
An ambitious Player vs. Player strategy game set on a dynamically sized 2D grid, where player-controlled units contend for scarce unit-construction resources in an effort to gain total control of the game map.

### Game Overview
At the beginning of the game, all players are allotted with a Mothership: a slow-moving capital unit which can attack enemy-controlled units and construct player-controlled Airship units. When a player's Mothership is destroyed, the player is defeated. Thus, the player's primary objective is to destroy all enemy Motherships.

The game map is a 2-Dimensional grid, dynamically sized to fit the number of participants. Each player Mothership is randomly placed on a valid tile (non-occupied, water tile). Once all Motherships have been placed, the game is begun.

Each round is time-limited and carried out synchronously for all players, meaning that the players are performing their turns at the same time. For each turn, every player-controlled unit has the opportunity to move and perform a single action. These actions include:
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

### Built With

* [Go](https://golang.org/)
* [AWS](https://aws.amazon.com/)
* [Ubuntu 18.04](http://releases.ubuntu.com/18.04.4/)
* [Docker](https://www.docker.com/)
* [Kubernetes](https://kubernetes.io/)
* [Terraform](https://www.terraform.io/)
* [go-sdl2](https://github.com/veandco/go-sdl2)


![Demo Image](https://cdn.discordapp.com/attachments/689340216284020812/689544753075060736/screenshot.png)
