# Snake [![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/tristangoossens/snake-go/game) [![Go Report Card](https://goreportcard.com/badge/github.com/tristangoossens/snake-go)](https://goreportcard.com/report/github.com/tristangoossens/snake-go) [![Build Status](https://travis-ci.com/tristangoossens/snake-go.svg?branch=master)](https://travis-ci.com/tristangoossens/snake-go)

This is a Terminal based snake game made by tristangoossens. The game is built by using the [termloop](https://github.com/JoelOtter/termloop) engine.

## The game

Here is a GIF of the first version of the game!

![alt text](https://github.com/tristangoossens/snake-go/blob/master/images/game-v1.gif "Version 1")

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

## Game settings

If you would like to alter some of the game settings you can do so by heading into the game directory and changing your preferences in the `gamesettings.go` file.

### How?

First of all, you will need to clone the repository into your workspace.

```shell
git clone https://github.com/tristangoossens/snake-go.git
```

After cloning you will have a **snake-go** directory in your workspace.

![alt text](https://github.com/tristangoossens/snake-go/tree/master/images/game-files.png "Snake game files")

Head into the game directory and open the `gameoptions.go` file to alter some settings to your desire.

![alt text](https://github.com/tristangoossens/snake-go/tree/master/images/gameoptions-file.png "Gameoptions in ")

After changing the settings to your desire, make sure to build the game.

```shell
go build run.go
```

Once the game is built you can run it using the following command!

```shell
./run
```

## Bucket list for future versions

- [x] Implementation of the skull mechanic
- [ ] Implement game difficulty options: easy, normal, hard
- [x] Add a setting panel for a user to change the color of the snake, food and border
- [x] Implement a restart button / quit button on gameoverscreen
- [x] Add score to Gameover screen
- [ ] Add instructions to the sidepanel
- [ ] Let user adjust arena size and snake speed(flags?)
- [ ] (could)Add function to check terminal size and adjust accordingly
- [ ] (could)Make a local database to save highscores.


