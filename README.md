## Snake [![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/tristangoossens/snake-go/game) [![Go Report Card](https://goreportcard.com/badge/github.com/tristangoossens/snake-go)](https://goreportcard.com/report/github.com/tristangoossens/snake-go) [![Build Status](https://travis-ci.com/tristangoossens/snake-go.svg?branch=master)](https://travis-ci.com/tristangoossens/snake-go)

![Logo](https://github.com/tristangoossens/snake-go/blob/master/images/snake-logo.png)  
This is a Terminal based snake game made by tristangoossens. ![GitHub followers](https://img.shields.io/github/followers/tristangoossens?style=social)  
Please star this repository to help my first big project grow! ![GitHub stars](https://img.shields.io/github/stars/tristangoossens/snake-go?style=social)  
**Documentation can be found at** https://github.com/tristangoossens/snake-go/tree/master/docs.

[![Run on Repl.it](https://repl.it/badge/github/tristangoossens/snake-go)](https://repl.it/github/tristangoossens/snake-go)  
[![Snake on itch.io](https://github.com/tristangoossens/snake-go/blob/master/images/itch-badge.png)](https://tristangoossens.itch.io/snake-go)

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

**There is a handful of ways to play snake**
 - Install the package
 - Install from itch (Windows) [![Snake on itch.io](https://github.com/tristangoossens/snake-go/blob/master/images/itch-badge.png)](https://tristangoossens.itch.io/snake-go)
 - Play on repl (online IDE) [![Run on Repl.it](https://repl.it/badge/github/tristangoossens/snake-go)](https://repl.it/github/tristangoossens/snake-go)
 - Clone the repository

### Install the package

First of all you will need Go, you can find more information [here](https://golang.org/).

When you have installed Go, you will need to install the game:

```shell
go get github.com/tristangoossens/snake-go
```

Then play it using the following command:

```shell
$GOPATH/bin/snake-go
```

### Cloning the repository

Cloning the repository is useful if you want to change any of the code or save your highscores to the markdown file.

How to clone:

```shell
git clone https://github.com/tristangoossens/snake-go.git
```

Then play it using:

```bash
go run run.go
```


## Bucket list for future versions

- [x] Implementation of the skull mechanic
- [x] Implement game difficulty options: easy, normal, hard
- [x] Add a setting panel for a user to change the color of the snake, food and border
- [x] Implement a restart button / quit button on gameoverscreen
- [x] Add score to Gameover screen
- [x] Add instructions to the sidepanel
- [x] (could)Create to save highscores. (markdown file)
- [x] Create a binary release
- [ ] Add test files for all game file
- [x] Rework title and gameover screens

## Links

- [Snake-go on termloop examples.](https://github.com/JoelOtter/termloop)
- [Snake-go on itch.io](https://tristangoossens.itch.io/snake-go)
- [Snake-go on github pages](https://tristangoossens.github.io/snake-go/)
- [Snake-go mentioned on golang weekly](https://golangweekly.com/issues/286)
