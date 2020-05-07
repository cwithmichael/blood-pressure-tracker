package com.cwithmichael.blood_pressure_tracker;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

@SpringBootApplication
@EnableJpaAuditing
public class BloodPressureTrackerApplication {

	public static void main(String[] args) {
		SpringApplication.run(BloodPressureTrackerApplication.class, args);
	}

}
