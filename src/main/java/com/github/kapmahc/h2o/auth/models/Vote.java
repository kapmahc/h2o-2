package com.github.kapmahc.h2o.auth.models;

import javax.persistence.*;
import java.io.Serializable;

@Entity
@Table(
        name = "votes",
        indexes = {
                @Index(name = "idx_votes_resource_type_id", columnList = "resourceType,resourceId", unique = true),
                @Index(name = "idx_votes_resource_type", columnList = "resourceType"),
        }
)
public class Vote implements Serializable {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    @Column(nullable = false)
    private String resourceType;
    @Column(nullable = false)
    private Long resourceId;
    @Column(nullable = false)
    private long point;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getResourceType() {
        return resourceType;
    }

    public void setResourceType(String resourceType) {
        this.resourceType = resourceType;
    }

    public Long getResourceId() {
        return resourceId;
    }

    public void setResourceId(Long resourceId) {
        this.resourceId = resourceId;
    }

    public long getPoint() {
        return point;
    }

    public void setPoint(long point) {
        this.point = point;
    }
}
