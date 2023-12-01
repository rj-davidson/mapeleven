import * as React from "react";
import Typography from "@mui/material/Typography";
import "../App.css";
import Grid from "@mui/material/Grid";
import { Tile } from "../Util/TileTS";
import MemberIDCard from "./MemberIDCard";
import jamesPhoto from "./james.jpg";
import ethanPhoto from "./ethan.jpg";
import rjPhoto from "./rj.jpg";
import ronniePhoto from "./ronnie.jpg";

function MemberPage() {
  return (
    <Grid container spacing={2}>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Tile
          sx={{ justifyContent: "left", alignItems: "left", textAlign: "left" }}
        >
          <Typography
            sx={{
              fontSize: { xs: "32px", sm: "32px" },
              marginLeft: "10px",
              fontWeight: "bold",
            }}
          >
            Team Members
          </Typography>
        </Tile>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Grid container spacing={2}>
          <Grid item xs={12} sm={12} md={12} lg={3} width="100%">
            <MemberIDCard
              name={"James Gibb"}
              age={20}
              contact={"jamesgibb27@gmail.com"}
              linkedin={"https://www.linkedin.com/in/jamesrobinsongibb/"}
              picture={jamesPhoto}
            ></MemberIDCard>
          </Grid>
          <Grid item xs={12} sm={12} md={12} lg={9} width="100%">
            <Tile
              sx={{
                justifyContent: "left",
                alignItems: "left",
                textAlign: "left",
                height: "100%",
              }}
            >
              <Typography
                sx={{
                  fontSize: { xs: "18px", sm: "20px" },
                  padding: "10px",
                }}
              >
                James is completing his B.S. in Data Science, and is set to
                graduate in December 2023. He has a passion for developing
                Machine Learning and Artificial Intelligence models, which led
                him to co-found Ratio, a startup focused on automating
                accounting and financial planning for small businesses. His role
                at Ratio has deepened his understanding of AI, especially with
                data analysis and management. Looking ahead, James aims to
                continue aiding business owners by automating their financial
                calculations and providing clear insights about their financial
                data. Outside of data science, he enjoys running, golfing, and
                is an avid sports fan supporting the University of Utah
                Football, the Cincinnati Bengals, and Manchester United.
              </Typography>
            </Tile>
          </Grid>
        </Grid>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Grid container spacing={2}>
          <Grid item xs={12} sm={12} md={12} lg={3} width="100%">
            <MemberIDCard
              name={"Ethan Quinlan"}
              age={20}
              contact={"contactethanquinlan@gmail.com"}
              linkedin={"https://www.linkedin.com/in/ethan-quinlan-a84210211/"}
              picture={ethanPhoto}
            ></MemberIDCard>
          </Grid>
          <Grid item xs={12} sm={12} md={12} lg={9} width="100%">
            <Tile
              sx={{
                justifyContent: "left",
                alignItems: "left",
                textAlign: "left",
                height: "100%",
              }}
            >
              <Typography
                sx={{
                  fontSize: { xs: "18px", sm: "20px" },
                  padding: "10px",
                }}
              >
                Ethan is currently enrolled as a BS/MS student at the University
                of Utah, with plans to graduate with his Masters in Computing in
                Spring 2025. His focus is on the Graphics and Visualizations
                specialization track for a Masters in Computing. This passion
                for graphics and visuals has proven invaluable in his role at
                MapEleven and during prior internships where he served as a
                front-end developer. MapEleven stands out as the pinnacle of his
                journey as a computer science student, showcasing the best of
                his abilities. In addition to his contributions at MapEleven, he
                co-created a graphical terrain generator/sculptor from the
                ground up. Looking ahead, he is eager to delve into back-end and
                full-stack development, aiming to broaden his expertise in web
                development. Beyond the realm of computer science, he finds joy
                in playing soccer and is a devoted fan of the Premier League,
                particularly Arsenal F.C.
              </Typography>
            </Tile>
          </Grid>
        </Grid>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Grid container spacing={2}>
          <Grid item xs={12} sm={12} md={12} lg={3} width="100%">
            <MemberIDCard
              name={"RJ Davidson"}
              age={20}
              contact={"hello@rjd.app"}
              linkedin={"https://www.linkedin.com/in/rj-davidson/"}
              picture={rjPhoto}
            ></MemberIDCard>
          </Grid>
          <Grid item xs={12} sm={12} md={12} lg={9} width="100%">
            <Tile
              sx={{
                justifyContent: "left",
                alignItems: "left",
                textAlign: "left",
                height: "100%",
              }}
            >
              <Typography
                sx={{
                  fontSize: { xs: "18px", sm: "20px" },
                  padding: "10px",
                }}
              >
                RJ is completing his B.S. in Software Development and will
                graduate in December 2023. He enjoys full-stack development,
                system architecture, and designing good user experiences—skills
                that he honed during his internship at Velvet Financial
                Services. Motivated by a mutual interest in cybersecurity, RJ,
                along with a group of like-minded peers, initiated the formation
                of the Cybersecurity Club at the University of Utah. He had the
                privilege of serving as president of the club for two semesters.
                Away from academia, he spends his time reading, favoring history
                and world-building series. He is also an avid supporter of
                several sports teams, including the Dallas Cowboys and Chelsea
                Football Club. With his sights set on the future, RJ aspires to
                engineer applications that value user privacy and security
                without sacrificing an intuitive user experience.
              </Typography>
            </Tile>
          </Grid>
        </Grid>
      </Grid>
      <Grid item xs={12} sm={12} md={12} lg={12} width="100%">
        <Grid container spacing={2}>
          <Grid item xs={12} sm={12} md={12} lg={3} width="100%">
            <MemberIDCard
              name={"Ronnie Koe"}
              age={20}
              contact={"ronnieckoe@gmail.com"}
              linkedin={"https://www.linkedin.com/in/ronnie-koe-780957260/"}
              picture={ronniePhoto}
            ></MemberIDCard>
          </Grid>
          <Grid item xs={12} sm={12} md={12} lg={9} width="100%">
            <Tile
              sx={{
                justifyContent: "left",
                alignItems: "left",
                textAlign: "left",
                height: "100%",
              }}
            >
              <Typography
                sx={{
                  fontSize: { xs: "18px", sm: "20px" },
                  padding: "10px",
                }}
              >
                Ronnie is currently pursuing a B.S. in Computer Science and will
                graduate in December 2023. He enjoys using artificial
                intelligence and machine learning to solve complex problems. He
                is currently conducting artificial intelligence related medical
                research in the ARM Lab run by Dr. Alan Kuntz at the University
                of Utah. He enjoys the challenges and learning that comes with
                research and enjoys collaborating with others that share the
                same passions. He plans to get an advanced degree to continue to
                do research in the artificial intelligence field. Apart from
                computer science, Ronnie has been on the University of Utah’s
                Call of Duty team for many years. Outside of university, Ronnie
                enjoys playing video games and watching sports and
                entertainment. A couple of the teams that Ronnie is an avid
                supporter of include the Stanley Cup Champion Vegas Golden
                Knights and FC Bayern Munich.
              </Typography>
            </Tile>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default MemberPage;
