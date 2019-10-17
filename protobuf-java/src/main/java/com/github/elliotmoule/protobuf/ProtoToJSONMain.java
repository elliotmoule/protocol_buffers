package com.github.elliotmoule.protobuf;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.util.JsonFormat;
import example.simple.Simple;

import java.util.Arrays;

public class ProtoToJSONMain {
    public static void main(String[] args) throws InvalidProtocolBufferException {

        Simple.SimpleMessage.Builder builder = Simple.SimpleMessage.newBuilder();

        // Builder returns itself, so can chain property setting together.

        // Simple Fields
        builder.setId(42)
                .setIsSimple(true)
                .setName("My Simple Message Name");

        // Repeated Fields
        builder.addSampleList(1)
                .addSampleList(2)
                .addSampleList(3)
                .addAllSampleList(Arrays.asList(4,5,6));

        // Print Protobuf as JSON
        String jsonString = JsonFormat.printer().print(builder);
        System.out.println(jsonString);

        // parse JSON into Protobuf.
        Simple.SimpleMessage.Builder builder2 = Simple.SimpleMessage.newBuilder();

        JsonFormat.parser()
                .ignoringUnknownFields()
                .merge(jsonString, builder2);

        System.out.println(builder2);
    }
}
