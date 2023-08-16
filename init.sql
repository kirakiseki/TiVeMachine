-- 1 database
DROP DATABASE IF EXISTS `tivemachine`;

CREATE DATABASE `tivemachine` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

USE `tivemachine`;

-- 2 tables
CREATE TABLE `user` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `description` varchar(1024) DEFAULT NULL,
  `sex` int(1) DEFAULT 0,
  CONSTRAINT `sex` CHECK (`sex` IN (0, 1))
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `program` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(1024) DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL,
  `category` varchar(255) DEFAULT NULL
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `channel` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(1024) DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE schedule (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `program_id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `start_time` bigint(20) NOT NULL,
  `end_time` bigint(20) NOT NULL,
  FOREIGN KEY (program_id) REFERENCES program(id) ON DELETE CASCADE,
  FOREIGN KEY (channel_id) REFERENCES channel(id) ON DELETE CASCADE,
  CONSTRAINT `start_time` CHECK (`start_time` < `end_time`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `subscription` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `schedule_id` int(11) NOT NULL,
  `alert_time` bigint(20) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
  FOREIGN KEY (schedule_id) REFERENCES schedule(id) ON DELETE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

-- 3 data
-- user
INSERT INTO
  `user` (`username`, `password`)
VALUES
  (
    'ishirai',
    'y1YB9mYC9+/gbANs79+cZgUGVeLTT0EB4w5h3xwYAgA='
  );

-- program
INSERT INTO
  `program` (`name`, `description`, `cover`, `category`)
VALUES
  (
    '新闻联播',
    '《新闻联播》是中央广播电视总台每日晚间播出的一档新闻节目，被称为“中国政坛的风向标”，节目宗旨为“宣传党和政府的声音，传播天下大事”。',
    'https://p3.img.cctvpic.com/photoAlbum/page/performance/img/2019/5/29/1559095198581_853.jpg',
    '新闻'
  ),
  (
    '焦点访谈',
    '《焦点访谈》栏目开办于1994年4月1日，由中央电视台新闻评论部创办，是以深度报道为主、以舆论监督见长的电视新闻评论性栏目。《焦点访谈》的舆论监督节目多年来为人们所关注和喜爱，选择“政府重视、群众关心、普遍存在”的选题，坚持“用事实说话”的方针，反映和推动解决了大量社会进步与发展过程中存在的问题。',
    'https://p1.img.cctvpic.com/photoAlbum/page/performance/img/2019/5/29/1559095229796_295.jpg',
    '新闻'
  ),
  (
    '今日说法',
    '《今日说法》诞生于1999年1月2日，是中央电视台第一档全日播法制栏目，已经成为家喻户晓的品牌栏目。栏目秉持“点滴记录中国法治进程”的理念，以“重在普法，监督执法，促进立法、服务百姓”为宗旨，全力打造“中国人的法律午餐”。栏目收视排名长期稳居央视前列，影响力持续增强；《12.4年度法治人物颁奖盛典》《小撒探会》等特别节目铸就高端品质，使栏目实现了更大的社会动员能力和普法功能。',
    'https://p5.img.cctvpic.com/photoAlbum/page/performance/img/2022/6/3/1654260241049_146.jpg',
    '法治'
  ),
  (
    '人与自然',
    '《人与自然》开播于1994年5月11日，栏目宗旨——“讴歌生命，关注环境。”栏目内容定位：介绍动植物和自然知识以及探索人与自然之间的相互影响，相互作用，探讨社会、经济、生态协调发展和可持续性发展的有效途径。',
    'https://p1.img.cctvpic.com/photoAlbum/page/performance/img/2019/5/20/1558324528137_138.jpg',
    '科教'
  ),
  (
    '晚间新闻',
    '《晚间新闻》是一档新闻品牌栏目，开办多年，一直受到观众喜欢。栏目信息量丰富，特色鲜明，风格清新，具有可视性、贴近性和权威性，是观众晚间了解当天重大新闻的窗口。',
    'https://p2.img.cctvpic.com/photoAlbum/page/performance/img/2023/3/27/1679907580454_154.jpg',
    '新闻'
  ),
  (
    '正大综艺',
    '文旅标杆，全新起航。2022年《正大综艺》全新出发，每期节目走进一座具有文旅特色的村镇，探访乡民心中的打卡胜地。 和《正大综艺》一起，饱览四方乡镇风光好物，细品历史今朝气象万千。',
    'https://p4.img.cctvpic.com/photoAlbum/page/performance/img/2022/4/24/1650791588920_696.jpg',
    '综艺'
  ),
  (
    '朝闻天下',
    '横跨CCTV-1综合、CCTV-13新闻2个频道，囊括3档新闻。日均累计收视率达2.2%。2006年央视原《早间新闻》改版为《朝闻天下》，时长增至180分钟。',
    'https://p3.img.cctvpic.com/photoAlbum/page/performance/img/2019/5/23/1558602307107_388.jpg',
    '新闻'
  ),
  (
    '新闻30分',
    '《新闻30分》是CCTV-13新闻频道与CCTV-1综合频道每天并机播出的一档新闻节目，其特点是时效性、高关注度和喜闻乐见，因其新颖、灵活的新闻报道直播方式，深受广大观众的好评和厚爱。',
    'https://p1.img.cctvpic.com/photoAlbum/page/performance/img/2018/3/1/1519868905156_753.jpg',
    '新闻'
  ),
  (
    '第一动画乐园',
    '《第一动画乐园》前身《动画城》，于1994年7月开播，是中央电视台也是全国第一个电视动画栏目，自开播起一直在CCTV-1次黄时段播出。栏目面向0-14岁孩子，是以推介国产优秀动画片以及国产动画新片为主的播出平台。栏目旨在打造动画乐园，带领孩子进入奇妙世界。',
    'https://p2.img.cctvpic.com/photoAlbum/page/performance/img/2023/3/28/1679966698854_541.jpg',
    '动画'
  ),
  (
    '中国文艺报道',
    '《中国文艺报道》以“文艺新力量，我们正发声”为口号，坚持繁荣文艺创作、推动文艺创新，并以全媒体传播新格局，创新内容传播形式，带领观众细品艺术百味，感受时代新风尚。',
    'https://p1.img.cctvpic.com/photoAlbum/page/performance/img/2020/9/29/1601367009594_169.jpg',
    '综艺'
  ),
  (
    '越战越勇',
    '《越战越勇》是由中央广播电视总台文艺节目中心制作播出的品牌栏目。2022年，节目改版为益智挑战类综艺。“为梦想携手组队，有你在能量加倍！”节目以趣味答题为主要载体，兼具娱乐性教育性，激励携手而来的搭档成就梦想，展现当下社会和谐人际关系百态图。',
    'https://p4.img.cctvpic.com/photoAlbum/page/performance/img/2023/2/14/1676339336083_237.jpg',
    '综艺'
  ),
  (
    '向幸福出发',
    '以“相亲交友”为载体，讲述青春奋斗故事，寻觅佳侣携手同行。“中科院心理研究所国民心理健康评估发展中心”的“婚恋匹配测验”技术支持嘉宾选拔。',
    'https://p5.img.cctvpic.com/photoAlbum/page/performance/img/2022/8/25/1661387936699_861.jpg',
    '综艺'
  ),
  (
    '文化十分',
    '《文化十分》栏目力图制作传播主流文化价值观传播阵地，用新闻视角解读文化、以文艺手法制作新闻；每期节目都将为观众及时发布权威政策、扫描梳理国内外重大文化事件、透视观察热点文化现象；主创团队还运用朴实的镜头讲述精品力作背后，艺术家难苦曲折、鲜为人知的艺术创作故事；并从特殊的视角，记录文艺工作者走进乡村、矿山、工厂、军营，深入生活、扎根人民采风创作的现场，反映他们用真情、真心、真诚的艺术创作奉献人民、回馈时代的感人事迹。',
    'https://p4.img.cctvpic.com/photoAlbum/page/performance/img/2019/5/21/1558407578099_446.jpg',
    '综艺'
  ),
  (
    '天天把歌唱',
    '《天天把歌唱》是CCTV-3综艺频道一档全新的音乐歌曲类节目，旨在弘扬传播中国优秀的音乐歌曲，让中国老百姓的精神和身心得以愉悦，歌手歌唱家的精彩演唱，成了该节目的一大亮点！',
    'https://p3.img.cctvpic.com/imagepic/ent/C30399/ibugu/images/img1314249259910694.jpg',
    '综艺'
  ),
  (
    '我的艺术清单',
    '中央广播电视总台央视综艺频道推出《我的艺术清单》,以最“接地气”的艺术视角和创新的节目呈现模式，为观众深度挖掘时代人物与艺术“跨界”相遇的正能量故事。',
    'https://p4.img.cctvpic.com/photoAlbum/page/performance/img/2020/2/21/1582277700875_931.jpg',
    '综艺'
  ),
  (
    '艺览天下',
    '艺览天下，览天下之艺术精华！综艺频道最新推出以介绍世界综艺舞台各种精彩表演为主的综艺节目，《艺览天下》！它囊括了世界范围内的顶尖艺术表演，如音乐舞蹈、马戏魔术、杂技技巧等，将向中国观众全面展示世界范围内的各种艺术发展和最新成就。',
    'https://p1.img.cctvpic.com/imagepic/ent/C21436/ibugu/images/img1311213989815977.jpg',
    '综艺'
  ),
  (
    '开门大吉',
    '《开门大吉》是中央电视台综艺频道推出的益智游戏类综艺节目，由尼格买提主持。节目鼓励普通人通过游戏闯关的方式实现自己的家庭梦想，通过多种艺术方法挖掘、展现普通人的人性光辉，和观众进行情感共鸣。',
    'https://p4.img.cctvpic.com/photoAlbum/page/performance/img/2013/12/20/1387544771140_581.jpg',
    '综艺'
  ),
  (
    '星光大道',
    '《星光大道》是CCTV-3综艺频道新推出的一档大型综艺栏目,该栏目一改以往娱乐节目以明星表演为主的局面，本着“百姓自娱自乐”的宗旨，突出大众参与性、娱乐性，力求为全国各地，各行各业的普通劳动者提供一个放声歌唱，展现自我的舞台。',
    'https://p1.img.cctvpic.com/photoAlbum/page/performance/img/2018/1/24/1516758411218_996.jpg',
    '综艺'
  ),
  (
    '回声嘹亮',
    'CCTV-3综艺频道《回声嘹亮》栏目以“向经典的文艺作品致敬”为核心，通过推荐人推荐自己心中的经典旋律和作品，讲述他们推荐的理由，用他们的声音唱响时代的旋律，表达不同人群对时代作品的不同情怀！',
    'https://p4.img.cctvpic.com/photoAlbum/page/performance/img/2019/5/23/1558602500237_142.jpg',
    '综艺'
  ),
  (
    '你好生活',
    '中央广播电视总台首档新青年文化公益旅行综艺节目。',
    'https://p4.img.cctvpic.com/photoAlbum/page/performance/img/2021/8/4/1628042960982_397.jpg',
    '综艺'
  );

-- channel
INSERT INTO
  `channel` (`name`, `description`, `cover`)
VALUES
  (
    'CCTV-1 综合',
    '中央电视台第一套节目综合频道，1958年开播。打造以新闻为主的精品综合频道，传递力量、温暖和希望。',
    'https://p1.img.cctvpic.com/photoAlbum/templet/common/DEPA1532314258547503/cctv-1_180817.png'
  ),
  (
    'CCTV-3 综艺',
    '中央电视台综艺频道是以播出音乐及歌舞节目为主的专业频道。',
    'https://p1.img.cctvpic.com/photoAlbum/templet/common/DEPA1532314258547503/cctv-3_180817.png'
  );

-- schedule
INSERT INTO
  `schedule` (
    `program_id`,
    `channel_id`,
    `start_time`,
    `end_time`
  )
VALUES
  -- 今日说法 00:09-00:37
  (3, 1, 1692115740, 1692117420),
  (3, 1, 1692202140, 1692203820),
  -- 人与自然 00:37-01:06
  (4, 1, 1692117420, 1692119160),
  (4, 1, 1692203820, 1692205560),
  -- 正大综艺 01:06-01：56
  (6, 1, 1692119160, 1692122160),
  (6, 1, 1692205560, 1692208560),
  -- 晚间新闻 01:56-02:26
  (5, 1, 1692122160, 1692123960),
  (5, 1, 1692208560, 1692210360),
  -- 今日说法 04:29-05:00
  (3, 1, 1692131340, 1692133200),
  (3, 1, 1692217740, 1692219600),
  -- 新闻联播 05:00-05:30
  (1, 1, 1692133200, 1692135000),
  (1, 1, 1692219600, 1692221400),
  -- 人与自然 05:30-06:00
  (4, 1, 1692135000, 1692136800),
  (4, 1, 1692221400, 1692223200),
  -- 朝闻天下 06:00-08:36
  (7, 1, 1692136800, 1692146160),
  (7, 1, 1692223200, 1692232560),
  -- 新闻30分 12:00-12:35
  (8, 1, 1692158400, 1692160500),
  (8, 1, 1692244800, 1692246900),
  -- 今日说法 12:35-13:11
  (3, 1, 1692160500, 1692162660),
  (3, 1, 1692246900, 1692249060),
  -- 第一动画乐园 17:14-18:17
  (9, 1, 1692177240, 1692181020),
  (9, 1, 1692263640, 1692267420),
  -- 新闻联播 19:00-19:39
  (1, 1, 1692183600, 1692185940),
  (1, 1, 1692270000, 1692272340),
  -- 焦点访谈 19:39-20:01
  (2, 1, 1692185940, 1692187260),
  (2, 1, 1692272340, 1692273660),
  -- 晚间新闻 22:00-22:34
  (5, 1, 1692194400, 1692196440),
  (5, 1, 1692280800, 1692282840),
  -- 我的艺术清单 01:11-02:02
  (15, 2, 1692119460, 1692122520),
  (15, 2, 1692205860, 1692208920),
  -- 艺览天下 02:02-02:54
  (16, 2, 1692122520, 1692125640),
  (16, 2, 1692208920, 1692212040),
  -- 中国文艺报道 04:22-04:43
  (10, 2, 1692130920, 1692132180),
  (10, 2, 1692217320, 1692218580),
  -- 越战越勇 04:43-06:01
  (11, 2, 1692132180, 1692136860),
  (11, 2, 1692218580, 1692223260),
  -- 向幸福出发 06:01-07:21
  (12, 2, 1692136860, 1692141660),
  (12, 2, 1692223260, 1692228060),
  -- 文化十分 07:21-07:32
  (13, 2, 1692141660, 1692142320),
  (13, 2, 1692228060, 1692228720),
  -- 天天把歌唱 07:32-0800
  (14, 2, 1692142320, 1692144000),
  (14, 2, 1692228720, 1692230400),
  -- 艺览天下 08:56-09:55
  (16, 2, 1692147360, 1692150900),
  (16, 2, 1692233760, 1692237300),
  -- 开门大吉 09:55-11:22
  (17, 2, 1692150900, 1692156120),
  (17, 2, 1692237300, 1692242520),
  -- 文化十分 11:22-11:34
  (13, 2, 1692156120, 1692156840),
  (13, 2, 1692242520, 1692243240),
  -- 星光大道 11:34-13:02
  (18, 2, 1692156840, 1692162120),
  (18, 2, 1692243240, 1692248520),
  -- 回声嘹亮 14:40-15:40
  (19, 2, 1692168000, 1692171600),
  (19, 2, 1692254400, 1692258000),
  -- 越战越勇 15:40-17:06
  (11, 2, 1692171600, 1692176760),
  (11, 2, 1692258000, 1692263160),
  -- 向幸福出发 19:30-20:53
  (12, 2, 1692185400, 1692190380),
  (12, 2, 1692271800, 1692276780);

-- 4 index
CREATE INDEX `idx_user_id` ON `user` (`id`);
CREATE INDEX `idx_program_id` ON `program` (`name`);
CREATE INDEX `idx_channel_id` ON `channel` (`name`);
CREATE INDEX `idx_schedule` ON `schedule` (`program_id`, `channel_id`, `start_time`, `end_time`);
