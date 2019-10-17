import simple_pb2 as simple_pb2

# Utilises reflection, so props and functions will not appear within intellisense.

simple_message = simple_pb2.SimpleMessage()

# Vars with default values are not included (boolean == False | int == 0)
simple_message.id = 123
simple_message.is_simple = True
simple_message.name = "This is a simple Message"
sample_list = simple_message.sample_list
sample_list.append(3)
sample_list.append(4)
sample_list.append(5)

print(simple_message)

with open("./protobuf-python/simple/simple.bin", "wb") as f:
    print("Write as binary")
    bytesAsString = simple_message.SerializeToString()
    f.write(bytesAsString)

with open("./protobuf-python/simple/simple.bin", "rb") as f:
    print("Read Values")
    simple_message_read = simple_pb2.SimpleMessage().FromString(f.read())

print(simple_message_read)

print("Read specific values")
print("IsSimple?: " + str(simple_message.is_simple))
print("ID: " + str(simple_message.id))