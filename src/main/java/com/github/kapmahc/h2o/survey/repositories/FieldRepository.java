package com.github.kapmahc.h2o.survey.repositories;

import com.github.kapmahc.h2o.survey.models.Field;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("survey.fieldR")
public interface FieldRepository extends CrudRepository<Field, Long> {
}
