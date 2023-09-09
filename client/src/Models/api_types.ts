enum SearchType {
  LEAGUE = "league",
  TEAM = "team",
  PLAYER = "player",
  FIXTURE = "fixture",
}

interface SearchModel {
  slug: string;
  name: string;
  type: SearchType;
  image: string;
}

interface CountryModel {
  code: string;
  name: string;
  flag: string;
}

enum LeagueType {
  CUP = "cup",
  LEAGUE = "league",
}

interface LeagueModel {
  slug: string;
  name: string;
  type: LeagueType;
  logo: string;
  country: CountryModel;
}