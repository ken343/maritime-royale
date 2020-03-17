<div align="center">
  
![GitHub contributors](https://img.shields.io/github/contributors/JosephZoeller/maritime-royale)
![GitHub forks](https://img.shields.io/github/forks/JosephZoeller/maritime-royale?label=Forks)
![GitHub stars](https://img.shields.io/github/stars/JosephZoeller/maritime-royale?style=Stars)
![GitHub issues](https://img.shields.io/github/issues-raw/JosephZoeller/maritime-royale)
[![Go Report Card](https://goreportcard.com/badge/github.com/JosephZoeller/maritime-royale)](https://goreportcard.com/report/github.com/JosephZoeller/maritime-royale)

</div>

<h3 align="center">
    
  ![destroyer_img](https://cdn.discordapp.com/attachments/689340216284020812/689360748920832000/destroyer.png) Maritime Royale ![destroyer_img](https://cdn.discordapp.com/attachments/689340216284020812/689360748920832000/destroyer.png)
  
A battle royale style, synchronous turn-based strategy game.
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
At the beginning of the game, all players are alotted with a Mothership: a slow-moving capital unit which can attack enemy-controlled units and construct player-controlled Airship units. When a player's Mothership is destroyed, the player is defeated. Thus, the player's primary objective is to destroy all enemy Motherships.

The game map is a 2-Dimensional grid, dynamically sized to fit the number of participants. Each player selects a valid tile (non-occupied, water) to initially place their Mothership. Once all Motherships have been placed, the game is begun. For each turn, every player-controlled unit has the opportunity to move and commit a single action.



### Built With

* [Go](https://golang.org/)
* [Docker](https://www.docker.com/)
* [Kubernetes](https://kubernetes.io/)
* [Terraform](https://www.terraform.io/)
* [go-sdl2](https://github.com/veandco/go-sdl2)


