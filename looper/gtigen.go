// Code generated by "goki generate -add-types"; DO NOT EDIT.

package looper

import (
	"goki.dev/gti"
)

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/looper.Ctr", IDName: "ctr", Doc: "Ctr combines an integer with a maximum value. It supports time tracking within looper.", Fields: []gti.Field{{Name: "Cur", Doc: "current counter value"}, {Name: "Max", Doc: "maximum counter value -- only used if > 0"}, {Name: "Inc", Doc: "increment per iteration"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/looper.Event", IDName: "event", Doc: "A Event has function(s) that can be called at a particular point\nin the loop, when the counter is AtCtr value.", Fields: []gti.Field{{Name: "Name", Doc: "Might be 'plus' or 'minus' for example."}, {Name: "AtCtr", Doc: "The counter value upon which this Event occurs."}, {Name: "OnEvent", Doc: "Callback function for the Event."}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/looper.NamedFunc", IDName: "named-func", Doc: "NamedFunc lets you keep an ordered map of functions.", Fields: []gti.Field{{Name: "Name"}, {Name: "Func"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/looper.NamedFuncs", IDName: "named-funcs", Doc: "NamedFunc is an ordered map of functions."})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/looper.NamedFuncsBool", IDName: "named-funcs-bool", Doc: "NamedFuncsBool is like NamedFuncs, but for functions that return a bool."})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/looper.Loop", IDName: "loop", Doc: "Loop contains one level of a multi-level iteration scheme.\nIt wraps around an inner loop recorded in a Stack, or around Main functions.\nIt records how many times the loop should be repeated in the Counter.\nIt records what happens at the beginning and end of each loop.\nFor example, a loop with 1 start, 1 end, and a Counter with max=3 will do:\nStart, Inner, End, Start, Inner, End, Start, Inner, End\nWhere the Inner loop is specified by a Stack or by Main,\nand Start and End are functions on the loop.\nSee Stack for more details on how loops are combined.", Fields: []gti.Field{{Name: "Counter", Doc: "Tracks time within the loop. Also tracks the maximum. OnStart, Main, and OnEnd will be called Ctr.Max times, or until IsDone is satisfied, whichever ends first."}, {Name: "OnStart", Doc: "OnStart is called at the beginning of each loop."}, {Name: "Main", Doc: "OnStart is called in the middle of each loop. In general, only use Main for the last Loop in a Stack. For example, actual Net updates might occur here."}, {Name: "OnEnd", Doc: "OnStart is called at the end of each loop."}, {Name: "IsDone", Doc: "If true, end loop. Maintained as an unordered map because they should not have side effects."}, {Name: "Events", Doc: "Events occur when Ctr.Cur gets to their AtCtr."}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/looper.Manager", IDName: "manager", Doc: "Manager holds data relating to multiple stacks of loops,\nas well as the logic for stepping through it.\nIt also holds helper methods for constructing the data.\nIt's also a control object for stepping through Stacks of Loops.\nIt holds data about how the flow is going.", Fields: []gti.Field{{Name: "Stacks", Doc: "map of stacks by Mode"}, {Name: "Mode", Doc: "The current evaluation mode."}, {Name: "isRunning", Doc: "Set to true while looping, false when done. Read only."}, {Name: "lastStartedCtr", Doc: "The Cur value of the Ctr associated with the last started level, for each timescale."}, {Name: "internalStop"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/looper.Stack", IDName: "stack", Doc: "Stack contains a list of Loops Ordered from top to bottom.\nFor example, a Stack might be created like this:\n\n\tmystack := manager.AddStack(etime.Train).AddTime(etime.Run, 2).AddTime(etime.Trial, 3)\n\tmyStack.Loops[etime.Run].OnStart.Add(\"NewRun\", initRunFunc)\n\tmyStack.Loops[etime.Trial].OnStart.Add(\"PresentTrial\", trialFunc)\n\nWhen run, myStack will behave like this:\ninitRunFunc, trialFunc, trialFunc, trialFunc, initRunFunc, trialFunc, trialFunc, trialFunc", Fields: []gti.Field{{Name: "Mode", Doc: "evaluation mode for this stack"}, {Name: "Loops", Doc: "An ordered map of Loops, from the outer loop at the start to the inner loop at the end."}, {Name: "Order", Doc: "The list and order of time scales looped over by this stack of loops,  ordered from top to bottom, so longer timescales like Run should be at the beginning and shorter timescales like Trial should be and the end."}, {Name: "StopNext", Doc: "If true, stop model at the end of the current StopLevel."}, {Name: "StopFlag", Doc: "If true, stop model ASAP."}, {Name: "StopLevel", Doc: "Time level to stop at the end of."}, {Name: "StopIterations", Doc: "How many iterations at StopLevel before actually stopping."}, {Name: "StepLevel", Doc: "Saved Time level for stepping -- what was set for last step or by gui."}, {Name: "StepIterations", Doc: "Saved number of steps for stepping -- what was set for last step or by gui."}}})
