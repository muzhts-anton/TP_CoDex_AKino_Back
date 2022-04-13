INSERT INTO
    users (
        username,
        password,
        email
    )
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

INSERT INTO
    feed (id, description, poster, PAGE, num)
VALUES
    (1, 'Топ 256', 'top.webp', 'collections', '1'),
    (
        2,
        'Приключения',
        'adventures.webp',
        'collections',
        2
    ),
    (
        3,
        'Для детей',
        'childish.webp',
        'collections',
        3
    ),
    (
        4,
        'По комиксам',
        'comics.webp',
        'collections',
        4
    ),
    (5, 'Драмы', 'drama.webp', 'collections', '5'),
    (
        6,
        'Для всей семьи',
        'family.webp',
        'collections',
        6
    ),
    (
        7,
        'Рекоммендации редакции',
        'ourTop.webp',
        'collections',
        7
    ),
    (
        8,
        'Романтические',
        'romantic.webp',
        'collections',
        8
    ),
    (
        9,
        'Спасение мира',
        'saveTheWorld.webp',
        'collections',
        9
    ),
    (
        10,
        'Советские',
        'soviet.webp',
        'collections',
        10
    ),
    (
        11,
        'Про шпионов',
        'spy.webp',
        'collections',
        11
    ),
    (12, 'Сказки', 'tales.webp', 'collections', 12);

INSERT INTO
    collections (id, title, description)
VALUES
    (
        1,
        'Топ 256',
        'Неповторимые must see'
    );

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
        genre,
        motto,
        director,
        budget,
        gross,
        duration,
        incollection
    )
VALUES
    (
        'showshenkRedemption.webp',
        'Побег из Шоушенка',
        'The Shawshank Redemption',
        9.0,
        1,
        '1994, США. Драма',
        'Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения.',
        'https://www.youtube.com/watch?v=PLl99DlL6b4',
        '1994',
        'США',
        'Драма',
        'Страх - это кандалы. Надежда - это свобода',
        'Франк Дарабонт',
        '$25 000 000',
        '$28 418 687',
        '142 мин.',
        1
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
        'Фантастика, Боевик, Приключения',
        '«Героями не рождаются. Героями становятся»',
        'Джон Фавро',
        '$140 000 000',
        '$585 366 247',
        '121 мин.',
        1
    );

INSERT INTO
    actors (
        imgsrc,
        name,
        nameoriginal,
        career,
        height,
        birthday,
        birthplace,
        genres,
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
        'Сказки',
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
        'Драма, комедия, короткометражка',
        259
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
    ratings (
        user_id,
        movie_id,
        rating
    )
VALUES
    (
        1,
        2,
        10
    );

INSERT INTO
    actors_actors (
        actor_id,
        relation_id
    )
VALUES
    (1, 2),
    (2, 1);

INSERT INTO
    movies_movies (
        movie_id,
        relation_id
    )
VALUES
    (1, 2),
    (2, 1);

INSERT INTO
    movies_actors (
        movie_id,
        actor_id
    )
VALUES
    (1, 1),
    (2, 2);
