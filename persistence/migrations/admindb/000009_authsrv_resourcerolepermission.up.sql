CREATE TABLE IF NOT EXISTS authsrv_resourcerolepermission (
    id integer NOT NULL,
    name character varying(256) NOT NULL,
    description character varying(512) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    modified_at timestamp with time zone NOT NULL,
    trash boolean NOT NULL,
    resource_permission_id integer NOT NULL,
    resource_role_id integer NOT NULL
);

ALTER TABLE authsrv_resourcerolepermission OWNER TO admindbuser;

CREATE SEQUENCE IF NOT EXISTS authsrv_resourcerolepermission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE authsrv_resourcerolepermission_id_seq OWNER TO admindbuser;

ALTER SEQUENCE authsrv_resourcerolepermission_id_seq OWNED BY authsrv_resourcerolepermission.id;

ALTER TABLE ONLY authsrv_resourcerolepermission ALTER COLUMN id SET DEFAULT nextval('authsrv_resourcerolepermission_id_seq'::regclass);

ALTER TABLE ONLY authsrv_resourcerolepermission ADD CONSTRAINT authsrv_resourcerolepermission_pkey PRIMARY KEY (id);

CREATE INDEX authsrv_resourcerolepermission_name_a65794e7 ON authsrv_resourcerolepermission USING btree (name);

CREATE INDEX authsrv_resourcerolepermission_name_a65794e7_like ON authsrv_resourcerolepermission USING btree (name varchar_pattern_ops);

CREATE INDEX authsrv_resourcerolepermission_resource_permission_id_c076e909 ON authsrv_resourcerolepermission USING btree (resource_permission_id);

CREATE INDEX authsrv_resourcerolepermission_resource_role_id_054f52d8 ON authsrv_resourcerolepermission USING btree (resource_role_id);

ALTER TABLE ONLY authsrv_resourcerolepermission
    ADD CONSTRAINT authsrv_resourcerole_resource_permission__c076e909_fk_authsrv_r FOREIGN KEY (resource_permission_id) 
    REFERENCES authsrv_resourcepermission(id) DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE ONLY authsrv_resourcerolepermission
    ADD CONSTRAINT authsrv_resourcerole_resource_role_id_054f52d8_fk_authsrv_r FOREIGN KEY (resource_role_id) 
    REFERENCES authsrv_resourcerole(id) DEFERRABLE INITIALLY DEFERRED;