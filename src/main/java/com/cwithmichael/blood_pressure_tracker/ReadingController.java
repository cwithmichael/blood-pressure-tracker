package com.cwithmichael.blood_pressure_tracker;

import java.util.List;

import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.CrossOrigin;


@CrossOrigin
@RestController
public class ReadingController {
    private final ReadingRepository repository;

    ReadingController(ReadingRepository repository) {
        this.repository = repository;
    }

    @GetMapping("/readings")
    List<Reading> all() {
        return repository.findAll();
    }

    @PostMapping("/readings")
    Reading newReading(@RequestBody Reading newReading) {
        return repository.save(newReading);
    }

    @GetMapping("/readings/{id}")
    Reading getReading(@PathVariable Long id) {
        return repository.findById(id)
            .orElseThrow(() -> new ReadingNotFoundException(id));
    }

    @PutMapping("/readings/{id}")
    Reading replaceReading(@RequestBody Reading newReading, @PathVariable Long id) {
  
      return repository.findById(id)
        .map(reading -> {
          reading.setSystolic(newReading.getSystolic());
          reading.setDiastolic(newReading.getDiastolic());
          reading.setPulse(newReading.getPulse());
          reading.setCreatedDate(newReading.getCreatedDate());
          return repository.save(newReading);
        })
        .orElseGet(() -> {
          newReading.setId(id);
          return repository.save(newReading);
        });
    }
  
    @DeleteMapping("/readings/{id}")
    void deleteReading(@PathVariable Long id) {
      repository.deleteById(id);
    }

}