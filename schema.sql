CREATE TABLE radicals(
   id                           INTEGER  NOT NULL PRIMARY KEY 
  ,object                       VARCHAR(7) NOT NULL
  ,url                          VARCHAR(38)
  ,data_updated_at              DATE 
  ,created_at                   DATE
  ,level                        INTEGER  NOT NULL
  ,slug                         VARCHAR(6) NOT NULL
  ,hidden_at                   DATE
  ,document_url                 VARCHAR(40)
  ,characters                   VARCHAR(1)
  ,meaning_mnemonic             VARCHAR(235) NOT NULL
  ,lesson_position              INTEGER  NOT NULL
  ,spaced_repetition_system_id  INTEGER  NOT NULL
);

CREATE TABLE meanings(
    sybject_id INTEGER NOT NULL
    ,meaning   TEXT 
    ,is_primary BOOLEAN CHECK (is_primary IN (0, 1))
    ,accepted_answer BOOLEAN CHECK (accepted_answer IN (0, 1))
);

CREATE TABLE auxiliary_meanings(
    sybject_id INTEGER NOT NULL
    ,meaning   TEXT 
    ,type TEXT
);

CREATE TABLE character_images(
    radical_id INTEGER NOT NULL
    ,url TEXT
    ,content_type TEXT
    ,inline_styles BOOLEAN
    ,color TEXT
    ,dimensions TEXT
    ,style_name TEXT
);

CREATE TABLE amalgamation_subject_ids(
    component_id integer NOT NULL
    ,set_id INTEGER NOT NULL
    ,unique (component_id, set_id)
);

CREATE TABLE kanji(
   id                          INTEGER  NOT NULL PRIMARY KEY 
  ,object                      VARCHAR(10) NOT NULL
  ,url                         VARCHAR(41)
  ,data_updated_at             DATE
  ,characters                  VARCHAR(1) 
  ,created_at                  DATE
  ,document_url                VARCHAR(40)
  ,hidden_at                   DATE
  ,lesson_position             INTEGER  NOT NULL
  ,level                       INTEGER  NOT NULL
  ,meaning_hint                VARCHAR(473) 
  ,meaning_mnemonic            VARCHAR(261) NOT NULL
  ,reading_mnemonic            VARCHAR(299) NOT NULL
  ,reading_hint                VARCHAR(509) 
  ,slug                        VARCHAR(1) NOT NULL
  ,spaced_repetition_system_id INTEGER  NOT NULL
);

CREATE TABLE readings(
   subject_id INTEGER NOT NULL
  ,type            VARCHAR(6)
  ,is_primary BOOLEAN  
  ,accepted_answer BOOLEAN NOT NULL
  ,reading         VARCHAR(2) NOT NULL
);


CREATE TABLE vocabularies(
   id                          INTEGER  NOT NULL PRIMARY KEY 
  ,object                      VARCHAR(10) NOT NULL
  ,url                         VARCHAR(41)
  ,data_updated_at             DATE
  ,characters                  VARCHAR(1) NOT NULL
  ,created_at                  DATE 
  ,document_url                VARCHAR(45)
  ,hidden_at                   DATE 
  ,lesson_position             INTEGER  NOT NULL
  ,level                       INTEGER  NOT NULL
  ,meaning_mnemonic            VARCHAR(171)
  ,reading_mnemonic            VARCHAR(411)
  ,slug                        VARCHAR(1) NOT NULL
  ,spaced_repetition_system_id BIGINT  NOT NULL
);

CREATE TABLE pronunciations(
  vocabulary_id INTEGER NOT NULL 
  ,url               VARCHAR(64) NOT NULL 
  ,gender            VARCHAR(4) NOT NULL
  ,source_id         INTEGER  NOT NULL
  ,pronunciation     VARCHAR(2) NOT NULL
  ,voice_actor_id    INTEGER  NOT NULL
  ,voice_actor_name  VARCHAR(7) 
  ,voice_description VARCHAR(12) 
  ,content_type      VARCHAR(9) NOT NULL
);

CREATE TABLE context_sentences(
  vocabulary_id INTEGER NOT NULL 
  ,en TEXT
  ,ja TEXT
);

CREATE TABLE parts_of_speech(
  vocabulary_id INTEGER NOT NULL 
  ,part_of_speech TEXT
);

CREATE TABLE assignments(
  id               INTEGER
  ,user_id         INTEGER  NOT NULL 
  ,object          VARCHAR(10) NOT NULL
  ,url             VARCHAR(48)
  ,data_updated_at DATE 
  ,created_at      DATE 
  ,subject_id      INTEGER  NOT NULL
  ,subject_type    VARCHAR(7) NOT NULL
  ,srs_stage       INTEGER  NOT NULL
  ,unlocked_at     DATE 
  ,started_at      DATE 
  ,passed_at       DATE 
  ,burned_at       DATE 
  ,available_at    DATE 
  ,resurrected_at  DATE 
  ,hidden          BOOLEAN
  ,is_learning     BOOLEAN NOT NULL DEFAULT false
  ,PRIMARY KEY (user_id,subject_id)
);
ALTER TABLE assignments 
ADD COLUMN is_learning BOOLEAN NOT NULL DEFAULT false;

CREATE TABLE spaced_repetition(
   id                        INTEGER  NOT NULL PRIMARY KEY 
  ,object                    VARCHAR(24) NOT NULL
  ,url                       VARCHAR(55)
  ,data_updated_at           DATE 
  ,created_at                DATE 
  ,name                      VARCHAR(38) NOT NULL
  ,description               VARCHAR(37) NOT NULL
  ,unlocking_stage_position  INTEGER  NOT NULL
  ,starting_stage_position   INTEGER  NOT NULL
  ,passing_stage_position    INTEGER  NOT NULL
  ,burning_stage_position    INTEGER  NOT NULL
);


CREATE TABLE srs_stages(
   srs_id        INTEGER  NOT NULL 
  ,interval      INTEGER
  ,position      INTEGER  NOT NULL
  ,interval_unit VARCHAR(30)
);

CREATE TABLE users(
  user_id INTEGER PRIMARY KEY NOT NULL
  ,level INTEGER NOT NULL
);
INSERT INTO users (user_id, level) VALUES (101, 4);

CREATE TABLE users_data(
  user_name VARCHAR(15) PRIMARY KEY NOT NULL
  ,user_id INTEGER NOT NULL
  ,pasword_sha VARCHAR(20) NOT NULL
);
INSERT INTO users_data (user_name,user_id, pasword_sha) VALUES ("zero", 101, "f9194e73f9e9459e3450ea10a179cdf77aafa695beecd3b9344a98d111622243");

CREATE TABLE sesion_ids(
  sesion_id VARCHAR(15) PRIMARY KEY NOT NULL
  ,user_id INTEGER NOT NULL
  ,created_at DATE NOT NULL
);