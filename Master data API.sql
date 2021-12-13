CREATE TABLE `teknisi` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(100),
  `specialist` varchar(100),
  `platform` varchar(100),
  `jumlah_antrian` int,
  `version` int
);

CREATE TABLE `kerusakan` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `jenis_kerusakan` varchar(100),
  `lama_pengerjaan` varchar(100),
  `harga` int,
  `version` int
);
