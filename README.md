# YeerChess

It's a pet project which implement the chess game invented by me when I was 13 years old.

## Demo & Screenshot

[http://yeerchess.github.io](http://yeerchess.github.io/)


![](http://content.screencast.com/users/yeer/folders/Jing/media/d7c7006e-fc70-4245-8e24-c6cdfc796c4f/00000001.png)

![](http://content.screencast.com/users/yeer/folders/Jing/media/2c55a7b4-7181-43b2-aadc-046bed753559/00000002.png)

![](http://content.screencast.com/users/yeer/folders/Jing/media/f469b6e9-16bd-48a3-87f5-eb79079730c0/00000003.png)


## How to Play?
Visit the [demo site](http://yeerchess.github.io/) and see the help section.


## Installation


```
	go get "github.com/athom/chess"
```


### Run it on the geek way


Go to the `cmd/yeerchessserver` direcotry

```
go install .
```

Go to the `cmd/yeerchessclient` diretory

```
go install .
```

Then start a server in your terminal

```
yeerchessserver
```

Start a client in the terminal

```
yeerchessclient
```

Start another client in the terminal

```
yeerchessclient
```

You can have fun in the console now.



### Run it as a web service.

First go to the `cmd/yeerchesswebserver` directory

```
go install .
```

Then start the server

```
yeerchesswebserver
```

It will start a webserver on port 3000.


Then check out another the HTML5 client [repo](https://github.com/yeerchess/yeerchess.github.io). 
Change [WS_HOST](https://github.com/yeerchess/yeerchess.github.io/blob/master/app/app.js#L421) in the file `app/app.js` to your localhost or any server you run the chess server.
Then open the index.html in your browser and you can see the App is playable. 


## How to play in the console mode?

The input use coordinate since mouse is not available. 
The format looks like this

```
x,y:m,n
```

The idea is `Point1:Point2`, means move piece from Point1 to Point2.

All possible poits on the board are:

```
(0,5) (1,5) (2,5) (3,5) (4,5) (5,5)

(0,4) (1,4) (2,4) (3,4) (4,4) (5,4)

(0,3) (1,3) (2,3) (3,3) (4,3) (5,3)

(0,2) (1,2) (2,2) (3,2) (4,2) (5,2)

(0,1) (1,1) (2,1) (3,1) (4,1) (5,1)

(0,0) (1,0) (2,0) (3,0) (4,0) (5,0)
```


So, for example, if you want to move your piece `3` from `(2,0)` to `(2,3)`, you just type

```
2,0:2,3
```


## License

MIT
