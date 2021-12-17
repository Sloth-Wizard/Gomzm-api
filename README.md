# Gomzm-api

main.go is used to test endpoints atm    

## Database

The database is constructed in time each time it encounters something that isn't in it      
It is constructed with a php backend at the moment

### Alliances table

This table contains alliances information

```sql
DROP TABLE IF EXISTS `alliances`;
CREATE TABLE IF NOT EXISTS `alliances` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creator_corporation_id` int NOT NULL,
  `creator_id` int NOT NULL,
  `date_founded` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `executor_corporation_id` int NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `ticker` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `alliance_id` int NOT NULL,
  `faction_id` int NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### Capsuleers table

This table contains the information of authorized capsuleer identified by SSO

```sql
DROP TABLE IF EXISTS `capsuleers`;
CREATE TABLE IF NOT EXISTS `capsuleers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `capsuleer_id` int NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `scopes` text NOT NULL,
  `datasource` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `token_type` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `refresh_token` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `owner` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `corporation_id` int NOT NULL,
  `access_token` text NOT NULL,
  `last_sso_call` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `access_token_exp` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `expires_in` int NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### Capsuleers public information table

This table contains public information about every capsuleers

```sql
DROP TABLE IF EXISTS `capsuleers_public`;
CREATE TABLE IF NOT EXISTS `capsuleers_public` (
  `id` int NOT NULL AUTO_INCREMENT,
  `alliance_id` int NOT NULL,
  `birthday` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `bloodline_id` int NOT NULL,
  `corporation_id` int NOT NULL,
  `description` text COLLATE utf8mb4_general_ci NOT NULL,
  `gender` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `race_id` int NOT NULL,
  `security_status` float(20,10) NOT NULL,
  `title` text COLLATE utf8mb4_general_ci NOT NULL,
  `character_id` int NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### Corporations table

This table contains corporations information

```sql
DROP TABLE IF EXISTS `corporations`;
CREATE TABLE IF NOT EXISTS `corporations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ceo_id` int NOT NULL,
  `creator_id` int NOT NULL,
  `description` text COLLATE utf8mb4_general_ci NOT NULL,
  `home_station_id` int NOT NULL,
  `member_count` int NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `shares` int NOT NULL,
  `tax_rate` float(20,10) NOT NULL,
  `ticker` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `corporation_id` int NOT NULL,
  `alliance_id` int NOT NULL,
  `date_founded` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `war_eligible` tinyint(1) NOT NULL,
  `faction_id` int NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### Factions table

This table contains factions information

```sql
DROP TABLE IF EXISTS `factions`;
CREATE TABLE IF NOT EXISTS `factions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `corporation_id` int NOT NULL,
  `description` text COLLATE utf8mb4_general_ci NOT NULL,
  `faction_id` int NOT NULL,
  `is_unique` tinyint(1) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `size_factor` float(20,10) NOT NULL,
  `solar_system_id` int NOT NULL,
  `station_count` int NOT NULL,
  `station_system_count` int NOT NULL,
  `militia_corporation_id` int NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### Items table

This table contains all the in game items information

```sql
DROP TABLE IF EXISTS `items`;
CREATE TABLE IF NOT EXISTS `items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `item_type_id` int NOT NULL,
  `capacity` float(20,10) NOT NULL,
  `description` text COLLATE utf8mb4_general_ci NOT NULL,
  `dogma_attributes` longtext COLLATE utf8mb4_general_ci NOT NULL,
  `dogma_effects` longtext COLLATE utf8mb4_general_ci NOT NULL,
  `group_id` int NOT NULL,
  `icon_id` int NOT NULL,
  `market_group_id` int NOT NULL,
  `mass` float(20,10) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `packaged_volume` float(20,10) NOT NULL,
  `portion_size` int NOT NULL,
  `published` tinyint(1) NOT NULL,
  `radius` float(20,10) NOT NULL,
  `type_id` int NOT NULL,
  `volume` float(20,10) NOT NULL,
  `graphic_id` int NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### Items flags table

This table contains where each item is fitted to a ship or if it's in a cargo/drone bay/container      
The data is fixed and imported from [fuzzwork-flags.csv](https://www.fuzzwork.co.uk/dump/latest/invFlags.csv)     
Please create the table and import the data after    

```sql
DROP TABLE IF EXISTS `items_flag`;
CREATE TABLE IF NOT EXISTS `items_flag` (
  `flag_id` int NOT NULL,
  `flag_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `flag_text` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `order_id` int NOT NULL
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### Killmails table

This table contains all a killmail details

```sql
DROP TABLE IF EXISTS `killmails`;
CREATE TABLE IF NOT EXISTS `killmails` (
  `id` int NOT NULL AUTO_INCREMENT,
  `killmail_id` int NOT NULL,
  `victim_character_id` int NOT NULL,
  `final_blow_character_id` int NOT NULL,
  `final_blow_faction_id` int NOT NULL,
  `killmail_details` longtext COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### Killmails hashes table

This table contains the hash of every killmails for verification and easier list building

```sql
DROP TABLE IF EXISTS `killmails_hash_id`;
CREATE TABLE IF NOT EXISTS `killmails_hash_id` (
  `id` int NOT NULL AUTO_INCREMENT,
  `killmail_hash` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `killmail_id` int NOT NULL,
  `corporation_id` int NOT NULL,
  `time` int NOT NULL,
  `capsuleer_id` int NOT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```
