package com.github.kapmahc.h2o.reading.repositories;

import com.github.kapmahc.h2o.reading.models.Book;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("reading.bookR")
public interface BookRepository extends JpaRepository<Book, Long> {
}
