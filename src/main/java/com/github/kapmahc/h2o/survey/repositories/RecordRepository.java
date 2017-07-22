package com.github.kapmahc.h2o.survey.repositories;

import com.github.kapmahc.h2o.survey.models.Record;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("survey.recordR")
public interface RecordRepository extends CrudRepository<Record, Long> {
}
