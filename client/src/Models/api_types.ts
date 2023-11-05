enum APISearchType {
    LEAGUE = 'league',
    TEAM = 'team',
    PLAYER = 'player',
    FIXTURE = 'fixture',
}

interface APISearch {
    slug: string;
    name: string;
    type: APISearchType;
    image: string;
}

interface APIScoreboard {
    date: APIScoreboardDate[];
}

interface APIScoreboardDate {
    date: string;
    leagues: {
        slug: string;
        name: string;
        type: string;
        logo: string;
        country?: APICountry;
        fixtures?: {
            slug: string;
            referee?: string;
            elapsed?: number;
            round?: number;
            status?: string;
            timezone?: string;
            homeTeam: {
                slug: string;
                name: string;
                logo: string;
            };
            awayTeam: {
                slug: string;
                name: string;
                logo: string;
            };
            score?: {
                home: number;
                away: number;
            };
        }[];
    }[];
}

interface APICountry {
    code: string;
    name: string;
    flag: string;
}

enum APILeagueType {
    CUP = 'cup',
    LEAGUE = 'league',
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
    logo: string;
    slug: string;
    name: APITeamName;
    leagueName: string;
    badge: string;
    founded?: number;
    nationalTeam?: boolean;
    country?: APICountry;
    competitions?: APICompetitions[];
    players?: APIPlayer[];
}

export interface APITeamStats {
    form: string;
    biggest: APITeamStatsBiggest;
    cards: APITSCards;
    clean_sheet: APITSResult;
    failed_to_score: APITSResult;
    fixtures: APITSFixtures;
    goals: APITSGoals;
    lineups: APITSFormation[];
    penalty: APITSPenalty;
}

export interface APITeamStatsBiggest {
    streak: APITSStreak;
    wins: APITSScore;
    loses: APITSScore;
    goals: APITSGoals;
}

export interface APITSStreak {
    wins: number;
    draws: number;
    loses: number;
}

export interface APITSScore {
    home: string;
    away: string;
}

export interface APITSGoals {
    for: APITSGoalDetail;
    against: APITSGoalDetail;
}

export interface APITSGoalDetail {
    home: number;
    away: number;
}

export interface APITSCards {
    yellow: APITSGoalMinuteSplit;
    red: APITSGoalMinuteSplit;
}

export interface APITSGoalMinuteSplit {
    '0-15': APITSGoalValueSplit;
    '16-30': APITSGoalValueSplit;
    '31-45': APITSGoalValueSplit;
    '46-60': APITSGoalValueSplit;
    '61-75': APITSGoalValueSplit;
    '76-90': APITSGoalValueSplit;
    '91-105': APITSGoalValueSplit;
    '106-120': APITSGoalValueSplit;
}

export interface APITSGoalValueSplit {
    total: number;
    percentage: string;
}

export interface APITSResult {
    home: number;
    away: number;
    total: number;
}

export interface APITSFixtures {
    played: APITSResult;
    wins: APITSResult;
    draws: APITSResult;
    loses: APITSResult;
}

export interface APITSFormation {
    formation: string;
    played: number;
}

export interface APITSPenalty {
    scored: APITSGoalValueSplit;
    missed: APITSGoalValueSplit;
    total: number;
}

interface APICompetitions {
    League: APILeague;
    Current: boolean;
    Stats?: APITeamStats;
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

interface APIPlayer {
    slug: string;
    name: string;
    firstname: string;
    lastname: string;
    age: number;
    height: string;
    weight: string;



    injured: boolean;
    photo: string;
}

interface APIFixture {
    slug: string;
    referee: string;
    timezone: string;
    date: Date;
    elapsed: number;
    round: number;
    status: string;
    homeTeamScore: number;
    awayTeamScore: number;
    season: APIFixtureSeason;
}

interface APIFixtureSeason {
    year: number;
    start: Date;
    end: Date;
    current: boolean;
}

interface APIFixtureTeam {
    slug: string;
    name: string;
    logo: string;
}

interface APIFixtureSet {
    fixture?: APIFixture;
    teams?: APIFixtureTeamDetails;
    events?: APIFixtureEvent[];
    lineups?: APIFixtureHomeAway;
}

interface APIFixtureHomeAway {
    home?: APIFixtureLineupInfo;
    away?: APIFixtureLineupInfo;
}

interface APIFixtureTeamDetails {
    home: APIFixtureTeamInfo;
    away: APIFixtureTeamInfo;
}

interface APIFixtureTeamInfo {
    slug: string;
    name: string;
    logo: string;
    winner?: boolean;
}

interface APIFixtureEvent {
    time: APIFixtureTimeDetail;
    team: APIFixtureTeamInfo;
    player: APIFixturePlayer;
    assist?: APIFixturePlayer;
    type: string;
    detail: string;
    comments?: string;
}

interface APIFixtureTimeDetail {
    elapsed: number;
    extra?: number;
}

interface APIFixturePlayer {
    slug: string;
    name: string;
}

interface APIFixtureLineupInfo {
    team: APIFixtureTeamInfo;
    coach: APIFixtureCoach;
    formation: string;
    matchPlayers: APIFixturePlayerDetail[];
}

interface APIFixtureCoach {
    name: string;
    photo: string;
}

interface APIFixturePlayerDetail {
    player: APIFixturePlayerDetailInfo;
}

interface APIFixturePlayerDetailInfo {
    slug: string;
    name: string;
    number: number;
    pos: string;
    x: number;
    y: number;
}


export {
    APIFixture,
    APIFixtureSet,
    APIFixtureEvent,
    APIFixtureTeam,
    APISearchType,
    APISearch,
    APIScoreboard,
    APIScoreboardDate,
    APICountry,
    APILeagueType,
    APILeague,
    APITeamName,
    APITeam,
    APICompetitions,
    APIStandings,
    APIRecord,
    APIStanding,
    APIPlayer,
};
