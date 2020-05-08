package com.cwithmichael.blood_pressure_tracker;

import lombok.Data;

import java.time.LocalDateTime;

import javax.persistence.Entity;
import javax.persistence.EntityListeners;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Column;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

@Data
@Entity
@EntityListeners(AuditingEntityListener.class)
public class Reading {
    private @Id @GeneratedValue Long id;
    @CreatedDate
    @Column(name = "reading_date")
    private LocalDateTime createdDate;
    private Integer systolic;
    private Integer diastolic;
    private Integer pulse;

    Reading () {}

    Reading(final Integer systolic, final Integer diastolic) {
        this.systolic = systolic;
        this.diastolic = diastolic;
    }

    Reading(final Integer systolic, final Integer diastolic, final Integer pulse) {
        this.systolic = systolic;
        this.diastolic = diastolic;
        this.pulse = pulse;
    }
}