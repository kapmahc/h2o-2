package com.github.kapmahc.h2o.reading.repositories;

import com.github.kapmahc.h2o.reading.models.Note;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("reading.noteR")
public interface NoteRepository extends JpaRepository<Note, Long> {
}
