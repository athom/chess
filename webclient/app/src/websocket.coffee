class Websocket
  #WS_HOST = "ws://localhost:3000"
  WS_HOST = "ws://localhost:7200"
  constructor: () ->
    @ws_conn = null


  connect: () ->
    if @ws_conn != null
      return

    @ws_conn = new WebSocket(WS_HOST + "/ws");
    @ws_conn.onopen = (data) ->
      console.log data

    @ws_conn.onmessage = (data) ->
      console.log data
      #debugger

    @ws_conn.onclose = (data) ->
      debugger

    @ws_conn.onerror = (data) ->
      alert "error"

  send: (data) ->
    if @ws_conn == null
      this.connect()

    @ws_conn.send(data)
