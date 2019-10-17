package com.github.elliotmoule.protobuf;

import example.simple.Simple.SimpleMessage;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Arrays;

public class SimpleMain {
    public static void main(String[] args) {
        System.out.println(("Hello World!"));

        SimpleMessage.Builder builder = SimpleMessage.newBuilder();

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

        // builder.clearSampleList();       // To Clear a List

        // builder.setSampleList(3, 42);    // At index 3, the value is 42.

        System.out.println(builder.toString());

        // Building the message allows for fetching the data.
        SimpleMessage message = builder.build();

        // Write the protocol buffers binary to file.
        try {
            FileOutputStream outputStream = new FileOutputStream("simple message.bin");

            message.writeTo(outputStream);
            outputStream.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

        // Send as byte array.
        // byte[] bytes = message.toByteArray();

        FileInputStream fileInputStream = null;
        try {
            System.out.println("Reading from file.");
            fileInputStream = new FileInputStream("simple message.bin");
            SimpleMessage messageFromFile = SimpleMessage.parseFrom(fileInputStream);
            System.out.println(messageFromFile);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

    }
}
