CREATE TABLE questions
(
    "id"       BIGSERIAL NOT NULL,
    question   text NOT NULL,
    is_active  boolean NOT NULL default true,
    created_at  date NOT NULL default now(),
    updated_at date NOT NULL,
    CONSTRAINT PK_9 PRIMARY KEY ( "id" )
);

CREATE TABLE choices
(
    "id"        BIGSERIAL NOT NULL,
    question_id BIGSERIAL NOT NULL,
    is_correct  boolean NOT NULL,
    choice      text NOT NULL,
    created_at  date NOT NULL default now(),
    updated_at  date NOT NULL,
    CONSTRAINT PK_15 PRIMARY KEY ( "id" ),
    CONSTRAINT FK_16 FOREIGN KEY ( question_id ) REFERENCES questions ( "id" )
);

CREATE INDEX FK_18 ON choices(
                              question_id
    );


CREATE TABLE user_question_choice
(
    "id"        bigint NOT NULL,
    user_id     bigint NOT NULL,
    choice_id   bigint NOT NULL,
    question_id bigint NOT NULL,
    is_correct  boolean NOT NULL,
    answer_time date NOT NULL default now(),
    CONSTRAINT PK_23 PRIMARY KEY ( "id" ),
    CONSTRAINT FK_24 FOREIGN KEY ( user_id ) REFERENCES users ( "id" ),
    CONSTRAINT FK_30 FOREIGN KEY ( question_id ) REFERENCES questions ( "id" ),
    CONSTRAINT FK_33 FOREIGN KEY ( choice_id ) REFERENCES choices ( "id" )
);

CREATE INDEX FK_26 ON user_question_choice (
                                            user_id
    );

CREATE INDEX FK_32 ON user_question_choice (
                                            question_id
    );

CREATE INDEX FK_35 ON user_question_choice (
                                            choice_id
    );
