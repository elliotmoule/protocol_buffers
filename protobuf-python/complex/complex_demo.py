import complex_pb2 as complex_pb2

complex_message = complex_pb2.ComplexMessage()

# The below is not allowed:
# ---------
# one_dummy_message = complex_pb2.DummyMessage()
# one_dummy_message.id = 123
# one_dummy_message.name = "I am a dummy message"
# complex_message.one_dummy = one_dummy_message
# ---------

complex_message.one_dummy.id = 123
complex_message.one_dummy.name = "I am a dummy message"

# One way to add to a list:
first_multiple_dummy = complex_message.multiple_dummy.add()
first_multiple_dummy.id = 345
first_multiple_dummy.name = "I'm the first element of an array"

# Another way to add to the list:
complex_message.multiple_dummy.add(id=567, name="Second element")

# When utilising 'extend', the inserted element becomes a copy (has it's own reference)
third_dummy_message = complex_pb2.DummyMessage()
third_dummy_message.id = 999
third_dummy_message.name = "I'm the last one!"
complex_message.multiple_dummy.extend([third_dummy_message])        

print(complex_message)