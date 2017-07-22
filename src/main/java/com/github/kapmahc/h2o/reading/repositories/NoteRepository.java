package com.github.kapmahc.h2o.reading.repositories;

import com.github.kapmahc.h2o.reading.models.Note;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("reading.noteR")
public interface NoteRepository extends CrudRepository<Note, Long> {
}
