package com.cwithmichael.blood_pressure_tracker;
import org.springframework.data.jpa.repository.JpaRepository;

interface ReadingRepository extends JpaRepository<Reading, Long> {
    
}