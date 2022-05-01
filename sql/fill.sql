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
    playlists (title, description, poster)
VALUES
    ('Топ 256', 'must see', 'top.webp'),
    ('Приключения', 'Захватывающий мир путешествий', 'adventures.webp'),
    ('Для детей', 'Самые маленькие оценят', 'childish.webp'),
    ('По комиксам', 'Экранизация культовых комиксов', 'comics.webp'),
    ('Драмы', 'Если захотелось поплакать', 'drama.webp'),
    ('Для всей семьи', 'В кургу самых близких', 'family.webp'),
    ('Рекоммендации редакции', 'С душой от AKino', 'ourTop.webp'),
    ('Романтические', 'Помечтаем?', 'romantic.webp'),
    ('Спасение мира', 'Мир в опасности!', 'saveTheWorld.webp'),
    ('Советские', 'Для старичков за 30', 'soviet.webp'),
    ('Про шпионов', 'Кажется, у нас завелась крыса', 'spy.webp'),
    ('Сказки', 'В тридевятом царстве, в тридесятом государстве...', 'tales.webp');

/* users playlists for tests and examples */
INSERT INTO
    playlists (title)
VALUES
    ('Ma boiz'),
    ('kinda trash');

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
    ),
    (
        'benedictCumberbatch.webp',
        'Бенедикт Камбербэтч',
        'Benedict Cumberbatch',
        'Актер, Продюсер',
        '183 см',
        '1976-07-19',
        'Лондон, Англия, Великобритания',
        184
    ),
    (
        'сhrisHemsworth.webp',
        'Крис Хемсворт',
        'Chris Hemsworth',
        'Актер, Продюсер',
        '190 см',
        '1983-08-11',
        'Мельбурн, Виктория, Австралия',
        125
    ),
    (
        'сhrisPratt.webp',
        'Крис Пратт',
        'Chris Pratt',
        'Актер, Продюсер',
        '188 см',
        '1979-06-21',
        'Вирджиния, Миннесота, США',
        155
    ),
    (
        'elizabethOlsen.webp',
        'Элизабет Олсен',
        'Elizabeth Olsen',
        'Актер, Продюсер',
        '168 см',
        '1989-02-16',
        'Шерман Оукс, Калифорния, США',
        77
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
    genres (genre, imgsrc, description, title)
VALUES
    ('action', 'Action.webp', 'Описание к action', 'Экшен'),
    ('adventure', 'Adventure.webp', 'Описание к adventure', 'Приключения'),
    ('anime', 'Anime.webp', 'Описание к anime', 'Аниме'),
    ('authors', 'Authors.webp', 'Описание к authors', 'Авторской'),
    ('biography', 'Biography.webp', 'Описание к biography', 'Биография'),
    ('cartoons', 'Cartoons.webp', 'Описание к cartoons', 'Мультфилмы'),
    ('comedy', 'Comedy.webp', 'Описание к comedy', 'Комедия'),
    ('criminal', 'Criminal.webp', 'Описание к criminal', 'Криминал'),
    ('detective', 'Detective.webp', 'Описание к detective', 'Деткктив'),
    ('documental', 'Documental.webp', 'Описание к documental', 'Документальное'),
    ('drama', 'Drama.webp', 'Описание к drama', 'Драма'),
    ('family', 'Family.webp', 'Описание к family', 'Для всей семьи'),
    ('fantasy', 'Fantasy.webp', 'Описание к fantasy', 'Фантастика'),
    ('historical', 'Historical.webp', 'Описание к historical', 'Историческое'),
    ('horror', 'Horror.webp', 'Описание к horror', 'Хоррор'),
    ('melodrama', 'Melodrama.webp', 'Описание к melodrama', 'Мелодрама'),
    ('musical', 'Musical.webp', 'Описание к musical', 'Мюзикл'),
    ('mystic', 'Mystic.webp', 'Описание к mystic', 'Мистика'),
    ('romantic', 'Romantic.webp', 'Описание к romantic', 'Романтика'),
    ('short', 'Short.webp', 'Описание к short', 'Короткометражное кино'),
    ('sport', 'Sport.webp', 'Описание к sport', 'Кино про спорт'),
    ('thriller', 'Thriller.webp', 'Описание к thriller', 'Триллер'),
    ('western', 'Western.webp', 'Описание к western', 'Вестерн');

INSERT INTO
    movies_genres (movie_id, genre)
VALUES
    (1, 'anime'),
    (1, 'action'),
    (1, 'mystic'),
    (2, 'anime'),
    (2, 'drama'),
    (2, 'horror'),
    (2, 'thriller'),
    (3, 'comedy'),
    (4, 'action'),
    (4, 'romantic'),
    (4, 'thriller'),
    (5, 'anime'),
    (5, 'fantasy'),
    (5, 'horror');

INSERT INTO
    actors_genres (actor_id, genre)
VALUES
    (1, 'anime'),
    (1, 'action'),
    (1, 'mystic'),
    (2, 'anime'),
    (2, 'action'),
    (2, 'horror'),
    (3, 'comedy'),
    (4, 'action'),
    (4, 'romantic'),
    (4, 'thriller'),
    (5, 'anime'),
    (5, 'fantasy'),
    (6, 'horror'),
    (7, 'fantasy');

INSERT INTO
    announced (poster, title, titleoriginal, info, description, trailer, releasedate, country, director)
VALUES
    (
        'doctorStrange.webp',
        'Доктор Стрэндж: В мультивселенной безумия',
        'Doctor Strange in the Multiverse of Madness',
        'Продолжение магических приключений Доктора Стрэнджа.',
        'Предстоящий американский супергеройский фильм, основанный на комиксах Marvel о Докторе Стрэндже, созданный Marvel Studios и распространяемый Walt Disney Studios Motion Pictures. Продолжение фильма «Доктор Стрэндж» (2016) и 28-я по счёту картина в медиафраншизе «Кинематографическая вселенная Marvel» (КВM).',
        'https://www.youtube.com/watch?v=aWzlQ2N6qqg',
        '2022-05-05',
        'США',
        'Сэм Рэйми'
    ),
    (
        'Thor4.webp',
        'Тор: Любовь и гром',
        'Thor: Love and Thunder',
        'Джейн Фостер берет на себя обязанности Бога-громовержца.',
        'Предстоящий американский супергеройский фильм, основанный на комиксах Marvel о Торе, созданный Marvel Studios и распространяемый Walt Disney Studios Motion Pictures. Прямое продолжение фильма «Тор: Рагнарёк» (2017) и 29-я по счёту картина в кинематографической вселенной Marvel (КВМ).',
        'https://www.youtube.com/watch?v=tgB1wUcmbbw',
        '2022-07-08 19:10:25-07',
        'США',
        'Тайка Вайтити'
    ),
    (
        'Guardians.webp',
        'Стражи Галактики. Часть 3',
        'Guardians of the Galaxy Vol. 3',
        'Третья часть приключений команды, защищающей галактику',
        'Предстоящий американский супергеройский фильм, основанный на комиксах Marvel о приключениях команды супергероев Стражей Галактики. Производством занимается Marvel Studios, а распространением — Walt Disney Studios Motion Pictures.',
        '-',
        '2023-05-04 19:10:25-07',
        'США',
        'Джеймс Ганн'
    );

INSERT INTO
    announced_actors (actor_id, announced_id)
VALUES
    (8, 1),
    (11, 1),
    (9, 2),
    (10, 3);

INSERT INTO
    announced_genres (announced_id, genre)
VALUES
    (1, 'fantasy'),
    (1, 'action'),
    (1, 'horror'),
    (1, 'adventure'),
    (2, 'adventure'),
    (2, 'fantasy'),
    (2, 'action'),
    (2, 'comedy'),
    (3, 'adventure'),
    (3, 'fantasy'),
    (3, 'action'),
    (3, 'comedy');

INSERT INTO
    announced_announced (announced_id, relation_id)
VALUES
    (1, 2),
    (1, 3),
    (2, 3),
    (3, 1),
    (3, 2);
