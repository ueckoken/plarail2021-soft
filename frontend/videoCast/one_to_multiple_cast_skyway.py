# WebRTCで1->複数に配信する
# 1人がsender

import asyncio
import websockets
import json
import ssl
import os
address = "0.0.0.0"
# cert = "C://Users/asika/OneDrive/ドキュメント/webRTC/vscode_live_server.cert.pem"
# key = "C://Users/asika/OneDrive/ドキュメント/webRTC/vscode_live_server.key.pem"


port = 8081
connection_num = 0
connections = []
sender_token = os.environ["SENDER_TOKEN"]  # ["127.0.0.1"]
print(sender_token)

rooms = {}
# keyが部屋id, valueが{"sender_socket", "peer_id", "connections"}
print("a")
# 後で通信している個人が本物か見分けるのもじっそうしないといけないきがする


# offerer == receiver
# answerer == sender

# idで判別して複数動画同時配信したい

async def server(websocket, path):
    global connection_num, connections, rooms
    remote_address = websocket.remote_address
    print(remote_address)
    connections.append(websocket)
    while True:
        # 受信
        try:
            received_packet = await websocket.recv()
        except:
            break
        dictionary = json.loads(received_packet)
        # msg_type, room_idで構成
        # connect_senderの時のみ+peer_id
        promises = []
        msg_type = dictionary["msg_type"]
        room_id = dictionary["room_id"]

        if room_id not in rooms:
            room = {
                "sender_socket": None,
                "peer_id": None,
                "connections": [websocket],
            }
            rooms[room_id] = room
        else:
            room = rooms[room_id]
            if websocket not in room["connections"]:
                room["connections"].append(websocket)
        # 現在の通信のwebsocketが入ったroom_idのroomが存在することを保証

        if msg_type == "connect_sender":
            if sender_token is None or sender_token == dictionary["sender_token"]:
                if room["sender_socket"] is None:
                    print("sender_connect")
                    room["sender_socket"] = websocket
                    # senderは上書き
                    room["peer_id"] = dictionary["peer_id"]
                    for connection in room["connections"]:
                        if connection is websocket:
                            break
                        print("send")
                        promise = connection.send(
                            json.dumps({
                                "msg_type": "connect_receiver",
                                "room_id": room_id,
                                "peer_id": room["peer_id"]
                            }))
                        promises.append(promise)
        elif msg_type == "connect_receiver":
            print("connect_receiver")
            if room["sender_socket"] is not None:
                print("send")
                promise = websocket.send(
                    json.dumps({
                        "msg_type": "connect_receiver",
                        "room_id": room_id,
                        "peer_id": room["peer_id"]
                    }))
                promises.append(promise)
        elif msg_type == "exit_room":
            if websocket in room["connections"]:
                room["connections"].remove(websocket)
            if room["sender_socket"] is websocket:
                room["sender_socket"] = None
                room["peer_id"] = None
        print("{}: {}".format(path, dictionary))
        for p in promises:
            a = await p

    # 接続が切れたらその接続を削除
    for room_id, room in rooms.items():
        print(room)
        if websocket in room["connections"]:
            room["connections"].remove(websocket)
        if room["sender_socket"] is websocket:
            room["sender_socket"] = None
            room["peer_id"] = None
    rooms = {room_id: room for room_id,
             room in rooms.items() if len(room["connections"]) != 0}
    # ルームに人がいなくなったら削除
    connections.remove(websocket)


#ssl_context = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
#ssl_context.load_cert_chain(cert, keyfile=key)

start_server = websockets.serve(server, address, port)  # , ssl=ssl_context)
# サーバー立ち上げ
asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
