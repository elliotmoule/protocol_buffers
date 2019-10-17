package com.github.elliotmoule.protobuf;

import example.complex.Complex;
import example.complex.Complex.*;

import java.util.Arrays;

public class ComplexMain {
    public static void main(String[] args) {
        System.out.println("Complex Example");

        DummyMessage oneDummy = newDummyMessage(55, "one dummy message");

        ComplexMessage.Builder builder = ComplexMessage.newBuilder();

        // Singular message field
        builder.setOneDummy(oneDummy);


        // Repeated field
        builder.addMultipleDummy(newDummyMessage(66, "second dummy"))
                .addMultipleDummy(newDummyMessage(67, "third dummy"))
                .addMultipleDummy(newDummyMessage(68, "fourth dummy"));

        builder.addAllMultipleDummy(Arrays.asList(
                newDummyMessage(69, "other dummy"),
                newDummyMessage(70, "other other dummy")
        ));

        ComplexMessage message = builder.build();

        System.out.println(message);

        // GET

        // message.getOneDummy();
        // message.getMultipleDummyList();
    }

    public static DummyMessage newDummyMessage(Integer id, String name) {
        DummyMessage.Builder dummyMessageBuilder = DummyMessage.newBuilder();
        DummyMessage message = dummyMessageBuilder.setName(name)
                .setId(id)
                .build();
        return message;
    }
}
