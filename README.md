# Objective

I am trying to create a simulation game. focusing on the economics on a galactic scale. The player controls an empire while other AI players will control the other empires. 
Each empire controls stars and the planets orbiting the stars. 

The basic economic unit is a 'pop' 

various goods will be produced, traded and consumed by the pops and various other entities.

The empires will compete for dominance. Through warfare, trade and industrial might.

There will be social dynamics too with pops promoting/demoting and getting angry/happy based on standard of living

# rambling below

it looks like I have forgotten what most of this code did. It seems I was looking into some kind of an ECS system.

I might instead focus on a market side of things for the moment.

so i am slightly more comfortable with the display code now. But i need to think through the entty structures so that i can switch between system and galaxy view.

what does a star represent?
actually two things. on system view, it is the center of attention. all planets are setup around it. 
on galaxy view it is a spot.

both need to be displayed using the same entity.
so we have a galaxy view version of star and a system version of star.

and that is just the display level. not even getting into other stuff.