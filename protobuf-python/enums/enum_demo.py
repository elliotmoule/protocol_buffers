import enum_example_pb2 as enum_example_pb2

enum_message = enum_example_pb2.EnumMessage()
enum_message.id = 345
enum_message.day_of_the_week = enum_example_pb2.THURSDAY

print(enum_message)

# Enums are represented by numerical value. The below demonstrates how this works.
# print(enum_message.day_of_the_week)
# print(enum_message.day_of_the_week == enum_example_pb2.THURSDAY)

with open("./protobuf-python/enums/enums.bin", "wb") as f:
    f.write(enum_message.SerializeToString())
    print("Wrote to a file")

with open("./protobuf-python/enums/enums.bin", "rb") as f:
    enum_message_read = enum_example_pb2.EnumMessage().FromString(f.read())
    print("Read a file")

print(enum_message_read)