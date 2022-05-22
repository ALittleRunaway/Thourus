
INSERT INTO thourus.company (name) VALUES ('Pied Piper');
INSERT INTO thourus.company (name) VALUES ('Hooli');
INSERT INTO thourus.company (name) VALUES ('Thourus');


INSERT INTO thourus.space (name, company_id) VALUES ('development', 1);
INSERT INTO thourus.space (name, company_id) VALUES ('sales', 1);
INSERT INTO thourus.space (name, company_id) VALUES ('development', 2);
INSERT INTO thourus.space (name, company_id) VALUES ('qa', 2);
INSERT INTO thourus.space (name, company_id) VALUES ('development', 3);
INSERT INTO thourus.space (name, company_id) VALUES ('management', 3);


INSERT INTO thourus.project (name, space_id) VALUES ('Pied Piper', 1);
INSERT INTO thourus.project (name, space_id) VALUES ('Nucleus', 3);
INSERT INTO thourus.project (name, space_id) VALUES ('Thourus', 5);


INSERT INTO thourus.role (name) VALUES ('admin');
INSERT INTO thourus.role (name) VALUES ('manager');
INSERT INTO thourus.role (name) VALUES ('employee');


INSERT INTO thourus.status (name) VALUES ('stable');
INSERT INTO thourus.status (name) VALUES ('pending');


INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Richard', 'Hendricks', 'richard_hendricks@gmail.com', 'richard.h', 'richard_pass', 1, 1, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Jared', 'Dunn', 'jared_dunn@gmail.com', 'jared.d', 'jared_pass', 2, 1, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Monica', 'Hall', 'monica_hall@gmail.com', 'monica.h', 'monica_pass', 2, 1, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Gilfoyle', 'Bertram', 'gilfoyle_bertram@gmail.com', 'gilfoyle.b', 'gilfoyle_pass', 3, 1, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Dinesh', 'Chugtai', 'dinesh_chugtai@gmail.com', 'dinesh.c', 'dinesh_pass', 3, 1, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Erlich', 'Bachman', 'erlich_bachman@gmail.com', 'erlich.b', 'erlich_pass', 3, 1, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Gavin', 'Belson', 'gavin_belson@gmail.com', 'gavin.b', 'gavin_pass', 1, 2, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Denpok', 'Singh', 'denpok_singh@gmail.com', 'denpok.s', 'denpok_pass', 2, 2, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Hoover', 'Clemens', 'hoover_clemens@gmail.com', 'hoover.c', 'hoover_pass', 3, 2, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Maria', 'Petrova', 'maria_petrova@gmail.com', 'maria.p', 'm', 1, 3, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Princess', 'Carolyn', 'princess_carolyn@gmail.com', 'princess.c', 'p', 2, 3, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Judah', 'Mannowdog', 'judah_mannowdog@gmail.com', 'judah.m', 'judah_pass', 2, 3, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('BoJack', 'Horseman', 'bojack_horseman@gmail.com', 'bojack.h', 'b', 3, 3, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Mr', 'Peanutbutter', 'mr_peanutbutter@gmail.com', 'mr.p', 'mr_pass', 3, 3, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Todd', 'Chavez', 'todd_chavez@gmail.com', 'todd.c', 'todd_pass', 3, 3, '2022-05-10 17:43:23');
INSERT INTO thourus.user (name, surname, email, login, password, role_id, company_id, registration_date) VALUES ('Diane', 'Nguyen', 'diane_nguyen@gmail.com', 'diane.h', 'diane_pass', 3, 3, '2022-05-10 17:43:23');

