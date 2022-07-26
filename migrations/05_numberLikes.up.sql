create table if not exists tb_number_likes (
    id varchar(36) not null,
    user_uid varchar(36) not null,
    post_pid varchar(36) not null,
    value_like boolean not null DEFAULT FALSE,
    created_at timestamp not null DEFAULT Now(),
    updated_at timestamp,
    deleted_at timestamp,
    constraint pk_number_likes primary key (id, user_uid, post_pid),
    constraint fk_pk_number_likes_0 foreign key (user_uid) references tb_user(id),
    constraint fk_pk_number_likes_1 foreign key (post_pid) references tb_post(id)
)