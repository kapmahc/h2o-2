package com.github.kapmahc.h2o.survey.repositories;

import com.github.kapmahc.h2o.survey.models.Field;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("survey.fieldR")
public interface FieldRepository extends JpaRepository<Field, Long> {
}
