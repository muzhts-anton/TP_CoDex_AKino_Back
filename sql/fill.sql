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
    ('Рекомендации редакции', 'С душой от AKino', 'ourTop.webp'),
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

-- INSERT INTO
--     movies (
--         id,
--         poster,
--         title,
--         titleoriginal,
--         rating,
--         votesnum,
--         info,
--         description,
--         trailer,
--         releaseyear,
--         country,
--         motto,
--         director,
--         budget,
--         gross,
--         duration
--     )
-- VALUES
--     (
--         1000001,
--         'showshenkRedemption.webp',
--         'Побег из Шоушенка',
--         'The Shawshank Redemption',
--         9.0,
--         1,
--         '1994, США. Драма.',
--         'Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения.',
--         'https://www.youtube.com/watch?v=PLl99DlL6b4',
--         '1994',
--         'США',
--         '«Страх - это кандалы. Надежда - это свобода»',
--         'Франк Дарабонт',
--         '$25 000 000',
--         '$28 418 687',
--         '142 мин.'
--     ),
--     (
--         1000002,
--         'ironman.webp',
--         'Железный Человек',
--         'Iron Man',
--         10.0,
--         5,
--         '2008, США, Канада. Фантастика, Боевик, Приключения.',
--         'Миллиардер-изобретатель Тони Старк попадает в плен к Афганским террористам, которые пытаются заставить его создать оружие массового поражения. В тайне от своих захватчиков Старк конструирует высокотехнологичную киберброню, которая помогает ему сбежать. Однако по возвращении в США он узнаёт, что в совете директоров его фирмы плетётся заговор, чреватый страшными последствиями. Используя своё последнее изобретение, Старк пытается решить проблемы своей компании радикально...',
--         'https://www.youtube.com/watch?v=PLl99DlL6b4',
--         '2008',
--         'США, Канада',
--         '«Героями не рождаются. Героями становятся»',
--         'Джон Фавро',
--         '$140 000 000',
--         '$585 366 247',
--         '121 мин.'
--     ),
--     (
--         1000003,
--         'greenMile.webp',
--         'Зеленая миля',
--         'The Green Mile',
--         9.1,
--         8,
--         '1999, США. Драма, Криминал.',
--         'Пол Эджкомб — начальник блока смертников в тюрьме «Холодная гора», каждый из узников которого однажды проходит «зеленую милю» по пути к месту казни. Пол повидал много заключённых и надзирателей за время работы. Однако гигант Джон Коффи, обвинённый в страшном преступлении, стал одним из самых необычных обитателей блока.',
--         'https://www.youtube.com/watch?v=PLl99DlL6b4',
--         '1994',
--         'США',
--         '«Пол Эджкомб не верил в чудеса. Пока не столкнулся с одним из них»',
--         'Фрэнк Дарабонт',
--         '$60 000 000',
--         '$286 801 374',
--         '189 мин.'
--     ),
--     (
--         1000004,
--         'shindlersList.webp',
--         'Список Шиндлера',
--         'The Green Mile',
--         8.8,
--         2,
--         '1993, США. Драма, биография, история, военный.',
--         'Фильм рассказывает реальную историю загадочного Оскара Шиндлера, члена нацистской партии, преуспевающего фабриканта, спасшего во время Второй мировой войны почти 1200 евреев.',
--         'https://www.youtube.com/watch?v=PLl99DlL6b4',
--         '1993',
--         'США',
--         '«Этот список - жизнь»',
--         'Стивен Зеллиан, Томас Кенилли',
--         '$22 000 000',
--         '$321 306 305',
--         '195 мин.'
--     ),
--     (
--         1000005,
--         'returnOfTheKing.webp',
--         'Властелин Колец',
--         'The Green Mile',
--         8.6,
--         3,
--         '2003, Новая Зеландия, США. Фэнтези, приключения, драма',
--         'Повелитель сил тьмы Саурон направляет свою бесчисленную армию под стены Минас-Тирита, крепости Последней Надежды. Он предвкушает близкую победу, но именно это мешает ему заметить две крохотные фигурки — хоббитов, приближающихся к Роковой Горе, где им предстоит уничтожить Кольцо Всевластья.',
--         'https://www.youtube.com/watch?v=PLl99DlL6b4',
--         '2003',
--         'Новая Зеландия, США',
--         '«There can be no triumph without loss. No victory without suffering. No freedom without sacrifice»',
--         'Питер Джексон',
--         '$94 000 000',
--         '$377 027 325',
--         '201 мин.'
--     );

INSERT INTO
    feed (playlist_id)
VALUES
    (1), (2), (3), (4), (5), (6), (7), (8), (9), (10), (11), (12);

-- INSERT INTO
--     playlists_movies (playlist_id, movie_id)
-- VALUES
--     (1, 1000001),
--     (1, 1000002),
--     (2, 1000003),
--     (2, 1000004),
--     (2, 1000005),
--     (13,1000001),
--     (13,1000002),
--     (14,1000001);

INSERT INTO
    users_playlists (user_id, playlist_id)
VALUES
    (2, 13),
    (2, 14);

INSERT INTO
    actors (
        id,
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
        1000001,
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
        1000002,
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
        1000003,
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
        1000004,
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
        1000005,
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
        1000006,
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
        1000007,
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
        1000008,
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
        1000009,
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
        1000010,
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
        1000011,
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
        1,1000001,'2022-04-10 15:47:24','good','Любимый фильм. Енто шыэдевр!'
    ),
    (
        1,
        1000003,
        '2022-04-10 15:47:24',
        'good',
        'Нормуль фильмец!'
    );

INSERT INTO
    ratings (user_id, movie_id, rating)
VALUES
    (1, 1000001, 8),
    (1, 1000002, 10);

INSERT INTO
    actors_actors (actor_id, relation_id)
VALUES
    (1000001, 1000002),
    (1000002, 1000003),
    (1000003, 1000004),
    (1000004, 1000005),
    (1000005, 1000006),
    (1000006, 1000007),
    (1000007, 1000001);

-- INSERT INTO
--     movies_movies (movie_id, relation_id)
-- VALUES
--     (1000001, 1000002),
--     (1000002, 1000003),
--     (1000003, 1000004),
--     (1000004, 1000005),
--     (1000005, 1000001);

-- INSERT INTO
--     movies_actors (movie_id, actor_id)
-- VALUES
--     (1000001, 1000001),
--     (1000001, 1000004),
--     (1000002, 1000002),
--     (1000003, 1000003),
--     (1000004, 1000005),
--     (1000005, 1000006),
--     (1000005, 1000007);

INSERT INTO
    genres (genre, imgsrc, description, title)
VALUES
    ('action', 'Action.webp', 'Описание к action', 'Экшен'),
    ('adventure', 'Adventure.webp', 'Описание к adventure', 'Приключения'),
    ('anime', 'Anime.webp', 'Описание к anime', 'Аниме'),
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
    ('short', 'Short.webp', 'Описание к short', 'Короткометражное кино'),
    ('sport', 'Sport.webp', 'Описание к sport', 'Кино про спорт'),
    ('thriller', 'Thriller.webp', 'Описание к thriller', 'Триллер'),
    ('western', 'Western.webp', 'Описание к western', 'Вестерн');

-- INSERT INTO
--     movies_genres (movie_id, genre)
-- VALUES
--     (1000001, 'anime'),
--     (1000001, 'action'),
--     (1000002, 'anime'),
--     (1000002, 'drama'),
--     (1000002, 'horror'),
--     (1000002, 'thriller'),
--     (1000003, 'comedy'),
--     (1000004, 'action'),
--     (1000004, 'thriller'),
--     (1000005, 'anime'),
--     (1000005, 'fantasy'),
--     (1000005, 'horror');

INSERT INTO
    actors_genres (actor_id, genre)
VALUES
    (1000001, 'anime'),
    (1000001, 'action'),
    (1000002, 'anime'),
    (1000002, 'action'),
    (1000002, 'horror'),
    (1000003, 'comedy'),
    (1000004, 'action'),
    (1000004, 'thriller'),
    (1000005, 'anime'),
    (1000005, 'fantasy'),
    (1000006, 'horror'),
    (1000007, 'fantasy');

INSERT INTO
    announced (poster, title, titleoriginal, info, description, trailer, releasedate, country, director)
VALUES
    -- (
    --     'doctorStrange.webp',
    --     'Доктор Стрэндж: В мультивселенной безумия',
    --     'Doctor Strange in the Multiverse of Madness',
    --     'Продолжение магических приключений Доктора Стрэнджа.',
    --     'Предстоящий американский супергеройский фильм, основанный на комиксах Marvel о Докторе Стрэндже, созданный Marvel Studios и распространяемый Walt Disney Studios Motion Pictures. Продолжение фильма «Доктор Стрэндж» (2016) и 28-я по счёту картина в медиафраншизе «Кинематографическая вселенная Marvel» (КВM).',
    --     'https://www.youtube.com/watch?v=aWzlQ2N6qqg',
    --     '2022-05-05',
    --     'США',
    --     'Сэм Рэйми'
    -- ),
    (
        'Thor4.webp',
        'Тор: Любовь и гром',
        'Thor: Love and Thunder',
        'Джейн Фостер берет на себя обязанности Бога-громовержца.',
        'Предстоящий американский супергеройский фильм, основанный на комиксах Marvel о Торе, созданный Marvel Studios и распространяемый Walt Disney Studios Motion Pictures. Прямое продолжение фильма «Тор: Рагнарёк» (2017) и 29-я по счёту картина в кинематографической вселенной Marvel (КВМ).',
        'https://www.youtube.com/watch?v=tgB1wUcmbbw',
        '2022-05-25',
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
        '2022-05-23',
        'США',
        'Джеймс Ганн'
    );

INSERT INTO
    announced_actors (actor_id, announced_id)
VALUES
    -- (8, 1),
    -- (11, 1),
    -- (1000009, 1),
    (3630, 1),
    -- (1000010, 2);
    (3613, 2);

INSERT INTO
    announced_genres (announced_id, genre)
VALUES
    -- (1, 'fantasy'),
    -- (1, 'action'),
    -- (1, 'horror'),
    -- (1, 'adventure'),
    (1, 'adventure'),
    (1, 'fantasy'),
    (1, 'action'),
    (1, 'comedy'),
    (2, 'adventure'),
    (2, 'fantasy'),
    (2, 'action'),
    (2, 'comedy');

INSERT INTO
    announced_announced (announced_id, relation_id)
VALUES
    -- (1, 2),
    -- (1, 3),
    (1, 2),
    -- (3, 1),
    (2, 1);

INSERT INTO 
    playlists_movies (playlist_id, movie_id)
SELECT
1, m.id
FROM  movies m
ORDER BY m.rating DESC
LIMIT 256;

-- Adventures
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (2, 685246),
    (2, 464963),
    (2, 4445150),
    (2, 841914),
    (2, 312),
    (2, 3498),
    (2, 476),
    (2, 401152),
    (2, 328),
    (2, 44745);

-- For kids
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (3, 4445150),
    (3, 841914),
    (3, 46483),
    (3, 476),
    (3, 401152),
    (3, 679486),
    (3, 775273),
    (3, 95232),
    (3, 279102),
    (3, 775276);

-- Comics
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (4, 61237),
    (4, 263531),
    (4, 822708),
    (4, 843650),
    (4, 843649),
    (4, 689066),
    (4, 822709),
    (4, 841263),
    (4, 462360),
    (4, 104904);

-- Drama
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (5, 61237),
    (5, 435),
    (5, 464963),
    (5, 404900),
    (5, 4445150),
    (5, 841914),
    (5, 258687),
    (5, 401152),
    (5, 361),
    (5, 679486);

-- For family
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (6, 4445150),
    (6, 535341),
    (6, 841914),
    (6, 46483),
    (6, 312),
    (6, 3498),
    (6, 476),
    (6, 401152),
    (6, 679486),
    (6, 427198);

-- Recomendations
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (7, 61237),
    (7, 435),
    (7, 685246),
    (7, 464963),
    (7, 404900),
    (7, 4445150),
    (7, 535341),
    (7, 841914),
    (7, 46483),
    (7, 258687);

-- Romantic
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (8, 464963),
    (8, 4445150),
    (8, 254776),
    (8, 958722),
    (8, 102130),
    (8, 257889),
    (8, 85182),
    (8, 1991),
    (8, 406366),
    (8, 427884);

-- Save the world
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (9, 61237),
    (9, 685246),
    (9, 258687),
    (9, 263531),
    (9, 822708),
    (9, 843650),
    (9, 843649),
    (9, 689066),
    (9, 2656),
    (9, 1091);

-- Soviet
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (10, 42664),
    (10, 25108),
    (10, 46483),
    (10, 77263),
    (10, 45549),
    (10, 42254),
    (10, 45612),
    (10, 45409),
    (10, 45283),
    (10, 43757);

-- SPY
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (11, 336434),
    (11, 408871),
    (11, 619),
    (11, 749540),
    (11, 462548),
    (11, 3394),
    (11, 8851),
    (11, 405627),
    (11, 840884),
    (11, 41381);

-- Fairy tales
INSERT INTO 
    playlists_movies (playlist_id, movie_id)
VALUES
    (12, 322),
    (12, 689),
    (12, 407636),
    (12, 688),
    (12, 8408),
    (12, 276762),
    (12, 48356),
    (12, 89515),
    (12, 841914),
    (12, 976642);

INSERT INTO
    movies_genres(movie_id, genre)
VALUES
    (370, 'anime'),
    (49684, 'anime'),
    (958722, 'anime'),
    (8221, 'anime'),
    (441, 'anime'),
    (280961, 'anime'),
    (963343, 'anime'),
    (1219417, 'anime'),
    (273209, 'anime'),
    (2428, 'anime'),
    (885317, 'anime'),
    (730665, 'anime'),
    (261127, 'anime');
