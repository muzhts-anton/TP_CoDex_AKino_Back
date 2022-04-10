INSERT INTO
    users (
        username,
        password,
        email,
        imgsrc
    )
VALUES
    (
        'Ванька',
        '$2a$10$XFUgsxdBN.UiILtfITr/4urH1WIQWBMkqvLnAgfiYpZNguvySCBAq', -- 1234abcd
        'ivan@mail.ru',
        '/profile.svg'
    );

INSERT INTO
    feed (id, description, poster, PAGE, num)
VALUES
    (1, 'Топ 256', 'top.png', 'collections', '1'),
    (
        2,
        'Приключения',
        'adventures.png',
        'collections',
        2
    ),
    (
        3,
        'Для детей',
        'childish.png',
        'collections',
        3
    ),
    (
        4,
        'По комиксам',
        'comics.png',
        'collections',
        4
    ),
    (5, 'Драмы', 'drama.png', 'collections', '5'),
    (
        6,
        'Для всей семьи',
        'family.png',
        'collections',
        6
    ),
    (
        7,
        'Рекоммендации редакции',
        'ourTop.png',
        'collections',
        7
    ),
    (
        8,
        'Романтические',
        'romantic.png',
        'collections',
        8
    ),
    (
        9,
        'Спасение мира',
        'saveTheWorld.png',
        'collections',
        9
    ),
    (
        10,
        'Советские',
        'soviet.png',
        'collections',
        10
    ),
    (
        11,
        'Про шпионов',
        'spy.png',
        'collections',
        11
    ),
    (12, 'Сказки', 'tales.png', 'collections', 12);

INSERT INTO
    collections (id, title, description)
VALUES
    (
        1,
        'Топ 256',
        'с любовью из бд'
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
        'showshenkRedemption.png',
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
        'hz.png',
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
        genres
    )
VALUES
    (
        'tales.png',
        'Баба Яга',
        'Baba Yaga',
        'Актёр',
        '160 см (без ступы)',
        'Неизвестно',
        'Русь',
        'Сказки'
    ),
    (
        'tales.png',
        'Роберт Дауни мл.',
        'Robert Downey Jr.',
        'Актер, Продюсер, Сценарист, Гений, Миллиардер, Плейбой, Филантроп.',
        '174 см (без костюма)',
        '1965-04-4',
        'Манхэттэн, Нью-Йорк, США',
        'Драма, комедия, короткометражка'
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
        current_date,
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
    (1, 2);

INSERT INTO
    movies_movies (
        movie_id,
        relation_id
    )
VALUES
    (1, 2);

INSERT INTO
    movies_actors (
        movie_id,
        actor_id
    )
VALUES
    (1, 1),
    (2, 2);

