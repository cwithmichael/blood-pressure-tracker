package com.cwithmichael.blood_pressure_tracker;

public class ReadingNotFoundException extends RuntimeException      {
    /**
     *
     */
    private static final long serialVersionUID = 3167768558397349695L;

    ReadingNotFoundException(Long id) {
        super("Could not find reading " + id);
    }
}