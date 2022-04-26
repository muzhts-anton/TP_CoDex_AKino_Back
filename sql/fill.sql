INSERT INTO
    users (username, password, email)
VALUES
    (
        'Ванька',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'ivan@vk.ru'
    ),
    (
        'tmp1',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'tmp1@vk.ru'
    ),
    (
        'tmp2',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'tmp2@vk.ru'
    ),
    (
        'tmp3',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'tmp3@vk.ru'
    ),
    (
        'tmp4',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'tmp4@vk.ru'
    );

/* static 'feed' */
INSERT INTO
    playlists (id, title, description, poster)
VALUES
    (1, 'Топ 256', 'must see', 'top.webp'),
    (2, 'Приключения', 'Захватывающий мир путешествий', 'adventures.webp'),
    (3, 'Для детей', 'Самые маленькие оценят', 'childish.webp'),
    (4, 'По комиксам', 'Экранизация культовых комиксов', 'comics.webp'),
    (5, 'Драмы', 'Если захотелось поплакать', 'drama.webp'),
    (6, 'Для всей семьи', 'В кургу самых близких', 'family.webp'),
    (7, 'Рекоммендации редакции', 'С душой от AKino', 'ourTop.webp'),
    (8, 'Романтические', 'Помечтаем?', 'romantic.webp'),
    (9, 'Спасение мира', 'Мир в опасности!', 'saveTheWorld.webp'),
    (10, 'Советские', 'Для старичков за 30', 'soviet.webp'),
    (11, 'Про шпионов', 'Кажется, у нас завелась крыса', 'spy.webp'),
    (12, 'Сказки', 'В тридевятом царстве, в тридесятом государстве...', 'tales.webp');

/* users playlists for tests and examples */
INSERT INTO
    playlists (id, title)
VALUES
    (13, 'Ma boiz'),
    (14, 'kinda trash');

INSERT INTO
    movies (
        poster,
        title,
        titleoriginal,
        rating,
        votesnum,
        info,
        description,
        trailer,
        releaseyear,
        country,
        motto,
        director,
        budget,
        gross,
        duration
    )
VALUES
    (
        'showshenkRedemption.webp',
        'Побег из Шоушенка',
        'The Shawshank Redemption',
        9.0,
        1,
        '1994, США. Драма.',
        'Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения.',
        'https://www.youtube.com/watch?v=PLl99DlL6b4',
        '1994',
        'США',
        '«Страх - это кандалы. Надежда - это свобода»',
        'Франк Дарабонт',
        '$25 000 000',
        '$28 418 687',
        '142 мин.'
    ),
    (
        'ironman.webp',
        'Железный Человек',
        'Iron Man',
        10.0,
        5,
        '2008, США, Канада. Фантастика, Боевик, Приключения.',
        'Миллиардер-изобретатель Тони Старк попадает в плен к Афганским террористам, которые пытаются заставить его создать оружие массового поражения. В тайне от своих захватчиков Старк конструирует высокотехнологичную киберброню, которая помогает ему сбежать. Однако по возвращении в США он узнаёт, что в совете директоров его фирмы плетётся заговор, чреватый страшными последствиями. Используя своё последнее изобретение, Старк пытается решить проблемы своей компании радикально...',
        'https://www.youtube.com/watch?v=PLl99DlL6b4',
        '2008',
        'США, Канада',
        '«Героями не рождаются. Героями становятся»',
        'Джон Фавро',
        '$140 000 000',
        '$585 366 247',
        '121 мин.'
    ),
    (
        'greenMile.webp',
        'Зеленая миля',
        'The Green Mile',
        9.1,
        8,
        '1999, США. Драма, Криминал.',
        'Пол Эджкомб — начальник блока смертников в тюрьме «Холодная гора», каждый из узников которого однажды проходит «зеленую милю» по пути к месту казни. Пол повидал много заключённых и надзирателей за время работы. Однако гигант Джон Коффи, обвинённый в страшном преступлении, стал одним из самых необычных обитателей блока.',
        'https://www.youtube.com/watch?v=PLl99DlL6b4',
        '1994',
        'США',
        '«Пол Эджкомб не верил в чудеса. Пока не столкнулся с одним из них»',
        'Фрэнк Дарабонт',
        '$60 000 000',
        '$286 801 374',
        '189 мин.'
    ),
    (
        'shindlersList.webp',
        'Список Шиндлера',
        'The Green Mile',
        8.8,
        2,
        '1993, США. Драма, биография, история, военный.',
        'Фильм рассказывает реальную историю загадочного Оскара Шиндлера, члена нацистской партии, преуспевающего фабриканта, спасшего во время Второй мировой войны почти 1200 евреев.',
        'https://www.youtube.com/watch?v=PLl99DlL6b4',
        '1993',
        'США',
        '«Этот список - жизнь»',
        'Стивен Зеллиан, Томас Кенилли',
        '$22 000 000',
        '$321 306 305',
        '195 мин.'
    ),
    (
        'returnOfTheKing.webp',
        'Властелин Колец',
        'The Green Mile',
        8.6,
        3,
        '2003, Новая Зеландия, США. Фэнтези, приключения, драма',
        'Повелитель сил тьмы Саурон направляет свою бесчисленную армию под стены Минас-Тирита, крепости Последней Надежды. Он предвкушает близкую победу, но именно это мешает ему заметить две крохотные фигурки — хоббитов, приближающихся к Роковой Горе, где им предстоит уничтожить Кольцо Всевластья.',
        'https://www.youtube.com/watch?v=PLl99DlL6b4',
        '2003',
        'Новая Зеландия, США',
        '«There can be no triumph without loss. No victory without suffering. No freedom without sacrifice»',
        'Питер Джексон',
        '$94 000 000',
        '$377 027 325',
        '201 мин.'
    );

INSERT INTO
    feed (playlist_id)
VALUES
    (1), (2), (3), (4), (5), (6), (7), (8), (9), (10), (11), (12);

INSERT INTO
    playlists_movies (playlist_id, movie_id)
VALUES
    (1, 1),
    (1, 2),
    (2, 3),
    (2, 4),
    (2, 5);

INSERT INTO
    users_playlists (user_id, playlist_id)
VALUES
    (1, 13),
    (1, 14);

INSERT INTO
    actors (
        imgsrc,
        name,
        nameoriginal,
        career,
        height,
        birthday,
        birthplace,
        total
    )
VALUES
    (
        'tales.webp',
        'Баба Яга',
        'Baba Yaga',
        'Актёр',
        '160 см (без ступы)',
        'Неизвестно',
        'Русь',
        500
    ),
    (
        'robertDowneyJr.webp',
        'Роберт Дауни мл.',
        'Robert Downey Jr.',
        'Актер, Продюсер, Сценарист, Гений, Миллиардер, Плейбой, Филантроп.',
        '174 см (без костюма)',
        '1965-04-04',
        'Манхэттэн, Нью-Йорк, США',
        259
    ),
    (
        'tomHanks.webp',
        'Том Хэнкс',
        'Tom Hanks',
        'Актер, Продюсер, Режиссер, Сценарист',
        '183 см',
        '1956-07-09',
        'Конкорд, Калифорния, США',
        399
    ),
    (
        'timRobbins.webp',
        'Тим Роббинс',
        'Tim Robbins',
        'Актер, Продюсер, Режиссер, Сценарист',
        '196 см',
        '1958-10-16',
        'Уэст-Ковина, штат Калифорния, США',
        213
    ),
    (
        'liamNeeson.webp',
        'Лиам Нисон',
        'Liam Neeson',
        'Актер, Продюсер',
        '193 см',
        '1958-07-07',
        'Беллимен, Северная Ирландия, Великобритания',
        302
    ),
    (
        'elijahWood.webp',
        'Элайджа Вуд',
        'Elijah Wood',
        'Актер, Продюсер, Режиссер',
        '168 см',
        '1981-01-28',
        'Сидар-Рапидс, Айова, США',
        271
    ),
    (
        'ianMcKellen.webp',
        'Иэн МакКеллен',
        'Ian McKellen',
        'Актер, Сценарист, Продюсер',
        '169 см',
        '1939-05-25',
        'Бернли, Ланкашир, Англия, Великобритания',
        295
    );

INSERT INTO
    comments (
        user_id,
        movie_id,
        commentdate,
        commenttype,
        content
    )
VALUES
    (
        1,
        1,
        '2022-04-10 15:47:24',
        'good',
        'Любимый фильм. Енто шыэдевр!'
    );

INSERT INTO
    ratings (user_id, movie_id, rating)
VALUES
    (1, 2, 10);

INSERT INTO
    actors_actors (actor_id, relation_id)
VALUES
    (1, 2),
    (2, 3),
    (3, 4),
    (4, 5),
    (5, 6),
    (6, 7),
    (7, 1);

INSERT INTO
    movies_movies (movie_id, relation_id)
VALUES
    (1, 2),
    (2, 3),
    (3, 4),
    (4, 5),
    (5, 1);

INSERT INTO
    movies_actors (movie_id, actor_id)
VALUES
    (1, 1),
    (1, 4),
    (2, 2),
    (3, 3),
    (4, 5),
    (5, 6),
    (5, 7);

INSERT INTO
    genres (genre)
VALUES
    ('tmp1'),
    ('tmp2'),
    ('tmp3'),
    ('tmp4'),
    ('tmp5'),
    ('tmp6'),
    ('tmp7'),
    ('tmp8'),
    ('tmp9');

INSERT INTO
    movies_genres (movie_id, genre)
VALUES
    (1, 'tmp1'),
    (1, 'tmp2'),
    (1, 'tmp7'),
    (2, 'tmp1'),
    (2, 'tmp4'),
    (2, 'tmp6'),
    (2, 'tmp9'),
    (3, 'tmp3'),
    (4, 'tmp2'),
    (4, 'tmp8'),
    (4, 'tmp9'),
    (5, 'tmp1'),
    (5, 'tmp5'),
    (5, 'tmp6');

INSERT INTO
    actors_genres (actor_id, genre)
VALUES
    (1, 'tmp1'),
    (1, 'tmp2'),
    (1, 'tmp7'),
    (2, 'tmp1'),
    (2, 'tmp2'),
    (2, 'tmp6'),
    (3, 'tmp3'),
    (4, 'tmp2'),
    (4, 'tmp8'),
    (4, 'tmp9'),
    (5, 'tmp1'),
    (5, 'tmp5'),
    (6, 'tmp6'),
    (7, 'tmp5');
