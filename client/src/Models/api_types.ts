enum APISearchType {
  LEAGUE = "league",
  TEAM = "team",
  PLAYER = "player",
  FIXTURE = "fixture",
}

interface APISearch {
  slug: string;
  name: string;
  type: APISearchType;
  image: string;
}

interface APICountry {
  code: string;
  name: string;
  flag: string;
}

enum APILeagueType {
  CUP = "cup",
  LEAGUE = "league",
}

interface APILeague {
  slug: string;
  name: string;
  type: APILeagueType;
  logo: string;
  country?: APICountry;
  leagueStandings?: APIStandings;
}

interface APITeamName {
  long: string;
  short: string;
}

interface APITeam {
  slug: string;
  name: APITeamName;
  badge: string;
  founded?: number;
  nationalTeam?: boolean;
  country?: APICountry;
  competitions?: APICompetitions[];
}

interface APICompetitions {
  League: APILeague;
  Current: boolean;
}

interface APIStandings {
  standings: APIStanding[];
}

interface APIRecord {
  Played: number;
  Won: number;
  Draw: number;
  Lost: number;
  GoalsFor: number;
  GoalsAgainst: number;
}

interface APIStanding {
  Rank: number;
  Description: string;
  League: APILeague;
  Team: APITeam;
  Points: number;
  GoalsDiff: number;
  Group: string;
  Form: string;
  Status: string;
  Home: APIRecord;
  Away: APIRecord;
  Overall: APIRecord;
}

export {
  APISearchType,
  APISearch,
  APICountry,
  APILeagueType,
  APILeague,
  APITeamName,
  APITeam,
  APICompetitions,
  APIStandings,
  APIRecord,
  APIStanding
};
