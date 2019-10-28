# Instructions

This markdown contains a handwritten guide on how to play snake-go!

*** 

## How to play the game

You can install the game in 2 different ways, you can get the game as a package, which will build the game and put the executable into the bin folder.  
You can also clone the repository, this will get you the entire repository and you will be able to use the save highscore function on the Game Over screen!

### 1. Get the game as a package

First of all you will need Go, you can find more information [here](https://golang.org/).

When you have installed Go, you will need to install the game:

```shell
go get -u github.com/tristangoossens/snake-go
```

Then play it using the following command:

```shell
$GOPATH/bin/snake-go
```

### 2.  Clone the game (use the save score option)

Alternatively if you want to get the game files and be able to save your scores you will need to clone the repository into your repository:

```shell
git clone https://github.com/tristangoossens/snake-go.git
```

Then build the game:

```shell
go build run.go
```

Once the game is built you can run it using the following command!

```shell
./run
```


## Game instructions

The Game is started when you press ENTER on the titlescreen:

![alt text](https://github.com/tristangoossens/snake-go/blob/master/images/game-titlescreen.png "Titlescreen")

Once you have started the game you will be able to move the snake using the arrow keys on your keyboard.

![alt text](https://github.com/tristangoossens/snake-go/blob/master/images/game-titlescreen.png "Game screen")




