package main

func Decision(c *Counter, o *Counter) bool {

	if len(c.history[o.id]) == 0 {
		var h hist
		h.req_val = o.value
		h.delta = o.value
		h.trust = 0.5
		h.total_req++
		h.accepted_req++
		h.decision = "yes"
		c.history[o.id] = append(c.history[o.id], h)
		return true
	} else {
		if (c.history[o.id][len(c.history[o.id])-1].trust / float64(o.value)) > 0.5 {
			var h hist
			h.req_val = o.value
			h.delta = o.value
			h.trust += 0.1
			h.total_req++
			h.accepted_req++
			h.decision = "yes"
			c.history[o.id] = append(c.history[o.id], h)
			return true
		} else {
			var h hist
			h.req_val = o.value
			h.delta = o.value
			h.trust -= 0.1
			h.total_req++
			h.decision = "no"
			c.history[o.id] = append(c.history[o.id], h)
			return true
		}
	}
}
