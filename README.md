# Snake [![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/tristangoossens/snake-go/game) [![Go Report Card](https://goreportcard.com/badge/github.com/tristangoossens/snake-go)](https://goreportcard.com/report/github.com/tristangoossens/snake-go) [![Build Status](https://travis-ci.com/tristangoossens/snake-go.svg?branch=master)](https://travis-ci.com/tristangoossens/snake-go) ![GitHub](https://img.shields.io/github/license/tristangoossens/snake-go)

![GitHub stars](https://img.shields.io/github/stars/tristangoossens/snake-go?style=social) ![GitHub followers](https://img.shields.io/github/followers/tristangoossens?style=social)

- [Snake-go on termloop examples.](https://github.com/JoelOtter/termloop)
- [Snake-go on itch.io](https://tristangoossens.itch.io/snake-go)
- [Snake-go on github pages](https://tristangoossens.github.io/snake-go/)

This is a Terminal based snake game made by tristangoossens. The game is built by using the [termloop](https://github.com/JoelOtter/termloop) engine.

## The game

Here is a GIF of the second version of the game! For version 1 [click here](https://github.com/tristangoossens/snake-go/tree/v1).

![GIF](https://github.com/tristangoossens/snake-go/blob/master/images/game-v2.gif)

## Game settings

Here is a demo on how the game settings panel works. for more information on this [click here](https://github.com/tristangoossens/snake-go/blob/master/docs/gameoptions.md).

![GIF](https://github.com/tristangoossens/snake-go/blob/master/images/gameoptions.gif)

## Save score

**IMPORTANT!**    
You need to clone the repository in order to use this function, for more information [click here](https://github.com/tristangoossens/snake-go/blob/master/docs/instructions.md).

![GIF](https://github.com/tristangoossens/snake-go/blob/master/images/savehighscore.gif)

## How to play

First of all you will need Go, you can find more information [here](https://golang.org/).

When you have installed Go, you will need to install the game:

```shell
go get github.com/tristangoossens/snake-go
```

Then play it using the following command:

```shell
$GOPATH/bin/snake-go
```

## Bucket list for future versions

- [x] Implementation of the skull mechanic
- [x] Implement game difficulty options: easy, normal, hard
- [x] Add a setting panel for a user to change the color of the snake, food and border
- [x] Implement a restart button / quit button on gameoverscreen
- [x] Add score to Gameover screen
- [x] Add instructions to the sidepanel
- [ ] Let user adjust arena size and snake speed(flags?)
- [ ] (could)Add function to check terminal size and adjust accordingly
- [x] (could)Make a local database to save highscores. (markdown file)
- [ ] Add test files for all game files


