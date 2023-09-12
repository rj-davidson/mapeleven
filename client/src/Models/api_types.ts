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
  country?: CountryModel;
  leagueStandings?: LeagueStandings;
}

interface TeamModel {
  slug: string;
  name: string;
  logo: string;
  country?: CountryModel;
  league?: LeagueModel;
}

interface LeagueStandings {
  standings: TeamStanding[];
}

interface TeamStanding {
  Team: TeamModel;
  League: LeagueModel;
  Position: number;
  PlayedGames: number;
  Won: number;
  Draw: number;
  Lost: number;
  Points: number;
  GoalsFor: number;
  GoalsAgainst: number;
  GoalDifference: number;
  Form: string;
}