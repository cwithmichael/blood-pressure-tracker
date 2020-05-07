package com.cwithmichael.blood_pressure_tracker;

import lombok.extern.slf4j.Slf4j;

import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
@Slf4j
public class LoadDatabase {

  @Bean
  CommandLineRunner initDatabase(ReadingRepository repository) {
    return args -> {
      log.info("Preloading " + repository.save(new Reading(120, 60, 60)));
      log.info("Preloading " + repository.save(new Reading(120, 60, 60)));
    };
  }
}