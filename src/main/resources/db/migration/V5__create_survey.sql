CREATE TABLE survey_models (
  id         BIGSERIAL PRIMARY KEY,
  title      VARCHAR(255)                NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  start_up   DATE                        NOT NULL,
  shut_down  DATE                        NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_survey_models_type
  ON survey_models (type);

CREATE TABLE survey_fields (
  id         BIGSERIAL PRIMARY KEY,
  label      VARCHAR(255)                NOT NULL,
  name       VARCHAR(255)                NOT NULL,
  value      VARCHAR(255)                NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(16)                 NOT NULL DEFAULT 'text',
  required   BOOLEAN                     NOT NULL DEFAULT TRUE,
  form_id    BIGINT REFERENCES survey_models,
  sort_order INT                         NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_survey_fields_name_form_id
  ON survey_fields (name, form_id);

CREATE TABLE survey_records (
  id         BIGSERIAL PRIMARY KEY,
  value      TEXT                        NOT NULL,
  form_id    BIGINT REFERENCES survey_models,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
